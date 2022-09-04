package main

import (
	"SuperStar/database"
	"SuperStar/models"
	"SuperStar/routes"
	"github.com/BurntSushi/toml"
	"log"
)

func main() {
	var config models.Config
	if _, err := toml.DecodeFile("conf.toml", &config); err != nil {
		panic(err)
	}
	database.ConnectDB(&config)
	app := routes.New()

	log.Fatal(app.Listen(":3000"))
}
