package handlers

import (
	"SuperStar/docs/common"
	"github.com/gofiber/fiber/v2"
)

// GetAboutByID 根据id获取个人主页
// @Summary 根据id获取个人主页
// @Description 根据id获取个人主页
// @Tags about
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.ResponseModel{data=[]models.User}
// @Router /api/v1/users/{id} [get]
func GetAboutByID(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(common.ResponseModel{
		Success: true,
		Message: "Success",
		Data:    id,
		Status:  fiber.StatusOK,
	})
}
