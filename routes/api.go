package routes

import (
	"SuperStar/handlers"
	"github.com/gofiber/fiber/v2"
)

func addAPIGroup(router *fiber.App, account *handlers.AccountHandler) {
	api := router.Group("/api")
	accountApi := api.Group("/account")
	{
		accountApi.Post("/add", account.CreateAccount())
	}
}
