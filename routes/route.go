package routes

import (
	_ "SuperStar/docs"
	"SuperStar/handlers"
	"SuperStar/internal/middlemare"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/google/wire"
)

var Provider = wire.NewSet(NewRoute)

func NewRoute(
	accountH *handlers.AccountHandler,
	loginH *handlers.LoginHandler,
	poemH *handlers.PoemHandler,
	tagH *handlers.TagHandler,
	circleH *handlers.CircleHandler,
	storageH *handlers.StorageHandler,
	jwtMiddle *middlemare.JwtMiddleware,
) *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(recover2.New())
	app.Use(logger.New(logger.Config{
		Format:     "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Shanghai",
	}))

	app.Get("/swagger/*", swagger.HandlerDefault)
	addAPIGroup(app, accountH, loginH, poemH, tagH, circleH, storageH, jwtMiddle)

	return app
}
