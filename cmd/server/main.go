package main

import (
	"fmt"

	"github.com/3lur/go-mall/cmd/server/wire"
	"github.com/3lur/go-mall/internal/common/data"
	"github.com/3lur/go-mall/pkg/config"
	"github.com/3lur/go-mall/pkg/http"
)

func main() {

	config := config.New()
	_, cleanup, _ := data.NewData(config)

	s, f, err := wire.NewApp(config)

	if err != nil {
		panic(err)
	}

	http.Run(s.ServerHTTP, fmt.Sprintf("%s:%d", config.App.Host, config.App.Port))

	defer f()
	defer cleanup()

}
