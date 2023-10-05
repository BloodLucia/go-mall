package configs

import "github.com/kelseyhightower/envconfig"

type App struct {
	Host string `required:"true" default:"127.0.0.1"`
	Port int    `required:"true" default:"3000"`
}

func AppStore() App {
	var app App
	envconfig.MustProcess("APP", &app)
	return app
}
