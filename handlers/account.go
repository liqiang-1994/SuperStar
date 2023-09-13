package handlers

import (
	"SuperStar/common"
	"SuperStar/internal/config"
	"SuperStar/internal/middlemare"
	"SuperStar/internal/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type AccountHandler struct {
	accountSrv *services.AccountService
	cfg        *config.Config
}

func NewAccountHandler(accountService *services.AccountService, cfg *config.Config) *AccountHandler {
	return &AccountHandler{accountSrv: accountService, cfg: cfg}
}

// account 个人主页
// @Summary 个人主页
// @Description 个人主页
// @Tags account
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Router /api/account/{id} [get]
func (s *AccountHandler) GetUserByID(c *fiber.Ctx) error {
	uid := c.Params("id")
	user := middlemare.ExtractTokenMetadata(c, s.cfg)
	self := false
	id, err := strconv.ParseInt(uid, 10, 64)
	if err != nil && user != nil && user.Id == id {
		self = true
	}
	result, err := s.accountSrv.Personal(c.UserContext(), uid, self)
	if err != nil {
		return c.JSON(common.Fail(err))
	}
	return c.JSON(common.Success(result))
}
