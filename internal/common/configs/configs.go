package configs

import (
	"fmt"

	"github.com/joho/godotenv"
)

type Configs struct {
	Database
}

func New(filename string) *Configs {
	if filename == "" {
		filename = ".env"
	}
	if err := godotenv.Load(filename); err != nil {
		panic(fmt.Errorf("failed to load [%s] file: %s", filename, err))
	}

	return &Configs{
		Database: DatabaseStore(),
	}
}
