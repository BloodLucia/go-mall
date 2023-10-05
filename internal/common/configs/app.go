package configs

import "github.com/kelseyhightower/envconfig"

type App struct {
	Host string `required:"true" default:"127.0.0.1"`
	Port int    `required:"true" default:"8000"`
	Env  string `default:"local"`
}

func AppStore() App {
	var app App
	envconfig.MustProcess("APP", &app)
	return app
}
