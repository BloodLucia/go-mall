package config

import (
	"fmt"

	"github.com/3lur/go-mall/internal/common/configs"
	"github.com/joho/godotenv"
)

type Config struct {
	configs.Database
}

func New(filename string) *Config {
	if filename == "" {
		filename = ".env"
	}
	if err := godotenv.Load(filename); err != nil {
		panic(fmt.Errorf("failed to load [%s] file: %s", filename, err))
	}

	return &Config{
		Database: configs.DatabaseStore(),
	}
}
