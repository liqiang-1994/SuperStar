package handlers

import (
	"SuperStar/common"
	"SuperStar/internal/model"
	"SuperStar/internal/services"
	"github.com/gofiber/fiber/v2"
)

type LoginHandler struct {
	accountSrv *services.AccountService
}

func NewLoginHandler(accountService *services.AccountService) *LoginHandler {
	return &LoginHandler{accountSrv: accountService}
}

// sendSms 发送短信验证码
// @Summary 发送短信验证码
// @Description 发送短信验证码
// @Tags login
// @Accept json
// @Produce json
// @Param param body object true "param"
// @Router /api/sendSms [post]
func (s *LoginHandler) SendSms(c *fiber.Ctx) error {
	loginReq := &model.LoginReq{}
	if err := c.BodyParser(loginReq); err != nil {
		return c.JSON(common.Fail(err))
	}
	existUser, _ := s.accountSrv.QueryByPhone(c.UserContext(), loginReq.UserName)
	var uid int64
	if existUser == nil {
		result, err := s.accountSrv.CreateAccount(c.UserContext(), loginReq)
		if err != nil {
			return c.JSON(common.Fail(err))
		}
		uid = result
	} else {
		uid = existUser.Id
	}

	key := string(rune(uid)) + "_" + loginReq.UserName
	_, err := s.accountSrv.SendSmsCode(key)
	if err != nil {
		return c.JSON(common.Fail(err))
	}

	return c.JSON(common.Success(loginReq))
}

// login 登录
// @Summary 登录
// @Description 登录
// @Tags login
// @Accept json
// @Produce json
// @Param param body object true "param"
// @Router /api/login [post]
func (s *LoginHandler) Login(c *fiber.Ctx) error {
	loginReq := &model.LoginReq{}
	if err := c.BodyParser(loginReq); err != nil {
		return c.JSON(common.Fail(err))
	}
	loginResp, err := s.accountSrv.Login(c.UserContext(), loginReq)
	if err != nil {
		return c.JSON(common.Fail(err))
	}
	return c.JSON(common.Success(loginResp))
}
