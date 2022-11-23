package routes

import (
	_ "SuperStar/docs"
	"SuperStar/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/google/wire"
)

var Provider = wire.NewSet(NewRoute)

func NewRoute(accountH *handlers.AccountHandler) *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format:     "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Shanghai",
	}))

	app.Get("/swagger/*", swagger.HandlerDefault)
	// Add api prefix
	//api := app.Group("/api")
	addAPIGroup(app, accountH)

	//AccountRouter(api, services.AccountService{})
	//accountApi := api.Group("/people")
	//accountApi.Post("/create")

	//v1User.Get("/:id", handlers.GetAboutByID)

	return app
}

//func NewRoute() *fiber.App {
//	app := fiber.New()
//	app.Use(cors.New())
//	app.Use(logger.New(logger.Config{
//		Format:     "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
//		TimeFormat: "2006-01-02 15:04:05",
//		TimeZone:   "Asia/Shanghai",
//	}))
//
//	app.Get("/swagger/*", swagger.HandlerDefault)
//	//addAPIGroup(app, accountH)
//	// Add api prefix
//	//api := app.Group("/api")
//	//
//	//accountApi := api.Group("/people")
//	//accountApi.Post("/create")
//
//	//v1User.Get("/:id", handlers.GetAboutByID)
//
//	return app
//}
