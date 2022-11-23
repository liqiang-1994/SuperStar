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
	//ctx, cancel := context.WithCancel(context.Background())
	return &App{
		config: conf,
		app:    app,
	}
}

//func newHttpServer(
//	conf *config.Config,
//	router *fiber.App,
//) *http.Server {
//	return &http.Server{
//		Addr:    ":" + conf.App.Port,
//		Handler: router,
//	}
//}

func main() {
	app, cleanup, err := InitApp()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	//app := routes.NewRoute()

	log.Fatal(app.app.Listen(":3000"))
}
