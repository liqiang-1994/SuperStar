package middlemare

import (
	"SuperStar/common"
	"SuperStar/internal/config"
	"SuperStar/internal/entity"
	"SuperStar/internal/model"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type JwtMiddleware struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewMiddleware(db *gorm.DB, cfg *config.Config) *JwtMiddleware {
	return &JwtMiddleware{db: db, cfg: cfg}
}

func (s *JwtMiddleware) JWTProtected(c *fiber.Ctx) error {
	var tokenStr string
	authorization := c.Get("Authorization")
	if strings.HasPrefix(authorization, "Bearer ") {
		tokenStr = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("token") != "" {
		tokenStr = c.Cookies("token")
	}
	if tokenStr == "" {
		return c.JSON(common.NotLogin("Not Login"))
	}

	tokenByte, err := jwt.Parse(tokenStr, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}
		return []byte(s.cfg.Auth.SecretKey), nil
	})
	if err != nil {
		return c.JSON(common.NotLogin(fmt.Sprintf("invalidate token: %v", err)))
	}
	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return c.JSON(common.NotLogin("invalid token claim"))
	}
	var user entity.User
	s.db.First(&user, "id = ?", claims["sub"])
	if strconv.FormatInt(user.Id, 10) != claims["sub"] {
		return c.JSON(common.NotLogin("invalid token"))
	}
	c.Locals("user", model.FilterUserRecord(&user))
	return c.Next()
}

func ExtractTokenMetadata(c *fiber.Ctx, cfg *config.Config) *model.UserResponse {
	var tokenStr string
	authorization := c.Get("Authorization")
	if strings.HasPrefix(authorization, "Bearer ") {
		tokenStr = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("token") != "" {
		tokenStr = c.Cookies("token")
	}
	if tokenStr == "" {
		return nil
	}
	tokenByte, err := jwt.Parse(tokenStr, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}
		return []byte(cfg.Auth.SecretKey), nil
	})
	if err != nil {
		return nil
	}
	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if ok && tokenByte.Valid {
		id, err := strconv.ParseInt(claims["sub"].(string), 10, 64)
		if err != nil {
			return nil
		}
		tokenUser := &model.UserResponse{
			Id: id,
		}
		return tokenUser
	}

	return nil
}
