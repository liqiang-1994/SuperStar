package models

type Config struct {
	DB mysql `toml:"mysql"`
}
type mysql struct {
	Url      string
	DBName   string
	Port     int
	UserName string
	Password string
}
