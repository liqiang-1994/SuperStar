package handlers

import (
	"SuperStar/docs/common"
	"SuperStar/internal/entity"
	"SuperStar/internal/services"
	"context"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type AccountHandler struct {
	accountSrv *services.AccountService
}

func NewAccountHandler(accountService *services.AccountService) *AccountHandler {
	return &AccountHandler{accountSrv: accountService}
}

func (s *AccountHandler) CreateAccount() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req entity.Account
		err := c.BodyParser(&req)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			var resp = common.Fail(err)
			return c.JSON(resp)
		}
		customContext, cancel := context.WithCancel(context.Background())
		defer cancel()
		result, err := s.accountSrv.CreateAccount(customContext, &req)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			var resp = common.Fail(err)
			return c.JSON(resp)
		}
		return c.JSON(result)
	}
}

//// GetUserByID 根据id获取个人信息
//// @Summary 根据id获取个人信息
//// @Description 根据id获取个人信息
//// @Tags about
//// @Accept json
//// @Produce json
//// @Param id path int true "User ID"
//// @Success 200 {object} models.ResponseModel{data=[]models.Account}
//// @Router /api/v1/users/{id} [get]
//func GetUserByID(c *fiber.Ctx) error {
//	id := c.Params("id")
//	return c.JSON(common.ResponseModel{
//		Success: true,
//		Message: "Success",
//		Data:    id,
//		Status:  fiber.StatusOK,
//	})
//}
//
//// ExistUserName 是否存在该昵称
//// @Summary 是否存在该昵称
//// @Description 是否存在该昵称
//// @Tags about
//// @Accept json
//// @Produce json
//// @Param userName path string true "Username"
//// @Success 200 {object} models.ResponseModel{data=[]models.Account}
//// @Router /api/v1/account/{userName} [get]
//func ExistUserName(c *fiber.Ctx) error {
//	id := c.Params("userName")
//	return c.JSON(common.ResponseModel{
//		Success: true,
//		Message: "Success",
//		Data:    id,
//		Status:  fiber.StatusOK,
//	})
//}
