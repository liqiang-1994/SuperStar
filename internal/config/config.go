package config

import (
	"github.com/BurntSushi/toml"
	"github.com/google/wire"
)

var Provider = wire.NewSet(NewCfg)

type Config struct {
	DB pg `toml:"postgres"`
}
type pg struct {
	Url      string
	DBName   string
	Port     int
	UserName string
	Password string
}

func NewCfg() (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile("conf.toml", &config); err != nil {
		panic(err)
	}
	return &config, nil
}
