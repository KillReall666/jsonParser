package config

import (
	"flag"

	"github.com/caarlos0/env"
)

type Config struct {
	Address      string `env:"RUN_ADDRESS"`
	DataFilePath string `env:"RUN_FILE_PATH"`
}

const (
	defaultServer       = "localhost:8080"
	defaultDataFilePath = "/app" //"/Users/valeriiamerkulova/golangProjects/jsonParser/ports.json"
)

func New() (*Config, error) {
	cfg := Config{}

	flag.StringVar(&cfg.Address, "a", defaultServer, "server address [host:port]")
	flag.StringVar(&cfg.DataFilePath, "d", defaultDataFilePath, "file path for json ports data")

	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
