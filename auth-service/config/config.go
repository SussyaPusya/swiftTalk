package config

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Postgres Postgres
}

type Postgres struct {
	User     string `env:"POSTGRES_USER" env-default:"postgres"`
	Password string `env:"POSTGRES_PASSWORD" env-default:"postgres"`
	Host     string `env:"POSTGRES_HOST" env-default:"localhost"`
	Port     int    `env:"POSTGRES_PORT" env-default:"5432"`
	Database string `env:"POSTGRES_DATABASE" env-default:"postgres"`
	MaxConn  int    `env:"POSTGRES_MAX_CONN" env-default:"10"`
	MinConn  int    `env:"POSTGRES_MIN_CONN" env-default:"2"`
}

func New() (*Config, error) {

	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Println(err)
	}

	fmt.Println(cfg)

	return &cfg, nil

}
