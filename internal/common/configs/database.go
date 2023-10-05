package configs

import "github.com/kelseyhightower/envconfig"

type Database struct {
	Driver   string `default:"mysql"`
	Host     string `required:"true" default:"127.0.0.1"`
	Port     int    `required:"true" default:"3306"`
	Username string `required:"true" default:"root"`
	Password string `required:"true"`
	DbName   string `required:"true"`
}

func DatabaseStore() Database {
	var db Database
	envconfig.MustProcess("DB", &db)
	return db
}
