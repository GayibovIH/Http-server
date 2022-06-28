package config

import (
	"github.com/caarlos0/env/v6"
	"log"
)

type Config struct {
	HttpPort int `env:"HTTP_PORT" envDefault:"8080"`
}

var cfg Config

func init() {
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}
}

func GetConfig() *Config {
	return &cfg
}
