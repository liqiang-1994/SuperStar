package main

import (
	"SuperStar/internal/config"
	"github.com/gofiber/fiber/v2"
	"log"
)

type App struct {
	config *config.Config
	app    *fiber.App
}

func NewApp(conf *config.Config, app *fiber.App) *App {
	return &App{
		config: conf,
		app:    app,
	}
}

// @title API
// @version 1.0
// @description This is an auto-generated API Docs for SuperStar.
// @contact.name API Support
// @contact.email luxyva@outlook.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	app, cleanup, err := InitApp()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	log.Fatal(app.app.Listen(":3000"))
}
