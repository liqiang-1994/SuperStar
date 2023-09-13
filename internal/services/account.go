package services

import (
	"SuperStar/internal/config"
	"SuperStar/internal/entity"
	"SuperStar/internal/model"
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type AccountRepo interface {
	CreateUser(ctx context.Context, u *entity.User) (int64, error)
	QueryById(ctx context.Context, id string) (*entity.User, error)
	QueryByPhone(ctx context.Context, phone string) (*entity.User, error)
}

type AccountService struct {
	accountRepo AccountRepo
	tm          Transaction
	rdb         *redis.Client
	cfg         *config.Config
}

func NewAccountService(account AccountRepo, tm Transaction, client *redis.Client, cfg *config.Config) *AccountService {
	return &AccountService{accountRepo: account, tm: tm, rdb: client, cfg: cfg}
}

func (s *AccountService) QueryByPhone(ctx context.Context, phone string) (*entity.User, error) {
	return s.accountRepo.QueryByPhone(ctx, phone)
}

func (s *AccountService) Personal(ctx context.Context, uid string, self bool) (*model.UserResponse, error) {
	user, err := s.accountRepo.QueryById(ctx, uid)
	if err != nil {
		return nil, err
	}
	resp := &model.UserResponse{}
	err = copier.Copy(resp, &user)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *AccountService) CreateAccount(ctx context.Context, m *model.LoginReq) (int64, error) {
	var (
		err error
		id  int64
	)
	u := &entity.User{
		Phone:      m.UserName,
		Status:     1,
		Follow:     0,
		Watch:      0,
		Up:         0,
		UserName:   "shiyou_" + strconv.FormatInt(time.Now().Unix(), 10),
		CreateTime: time.Now(),
	}
	err = s.tm.ExecTx(ctx, func(ctx context.Context) error {
		id, err = s.accountRepo.CreateUser(ctx, u)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *AccountService) SendSmsCode(u string) (string, error) {
	return s.rdb.Set(u, "1234", 15*time.Minute).Result()
}

func (s *AccountService) Login(ctx context.Context, m *model.LoginReq) (*model.LoginResp, error) {
	var (
		err       error
		loginResp model.LoginResp
		uid       int64
	)
	u, err := s.accountRepo.QueryByPhone(ctx, m.UserName)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if u == nil {
		uid, err = s.CreateAccount(ctx, m)
		if err != nil {
			return nil, err
		}
	} else {
		uid = u.Id
	}
	//key := strconv.FormatInt(uid, 10) + "_" + m.UserName
	code := "1234"
	//code, err := s.rdb.Get(key).Result()
	//if err != nil {
	//	return nil, errors.New("验证码已过期")
	//}
	if code != m.CheckCode {
		return nil, errors.New("验证码错误")
	}
	s.rdb.HSetNX("user", m.UserName, uid)
	tokenByte := jwt.New(jwt.SigningMethodHS256)
	now := time.Now().UTC()
	claims := tokenByte.Claims.(jwt.MapClaims)
	claims["sub"] = strconv.FormatInt(uid, 10)
	claims["exp"] = now.Add(time.Hour * 24 * 7).Unix()
	claims["iat"] = now.Unix()
	claims["iss"] = "SuperStar"
	claims["nbf"] = now.Unix()
	tokenStr, err := tokenByte.SignedString([]byte(s.cfg.Auth.SecretKey))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("生成token失败: %v", err))
	}
	loginResp.Token = "Bearer " + tokenStr
	return &loginResp, nil
}
