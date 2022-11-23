//go:build wireinject
// +build wireinject

package main

import (
	"SuperStar/handlers"
	"SuperStar/internal/config"
	"SuperStar/internal/repos"
	"SuperStar/internal/services"
	"SuperStar/routes"
	"github.com/google/wire"
)

func InitApp() (*App, func(), error) {
	panic(wire.Build(config.Provider, repos.Provider, services.Provider, handlers.Provider, routes.Provider, NewApp))
}
