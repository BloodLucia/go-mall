package config

import (
	"flag"

	"github.com/3lur/go-mall/internal/common/configs"
	"github.com/3lur/go-mall/pkg/console"
	"github.com/joho/godotenv"
)

type Config struct {
	Database configs.Database
	App      configs.App
}

func New() *Config {
	var filename string

	// register command, usage: --config=.env
	flag.StringVar(&filename, "config", ".env", "config file, eg: --config=[.filename]")
	flag.Parse()

	if filename == "" {
		filename = ".env"
	}
	if err := godotenv.Load(filename); err != nil {
		console.ExitIf(err)
	}

	return &Config{
		Database: configs.DatabaseStore(),
		App:      configs.AppStore(),
	}
}
