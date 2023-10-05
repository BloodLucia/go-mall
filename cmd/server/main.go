package main

import (
	mallcmd "github.com/3lur/go-mall/cmd"
	"github.com/3lur/go-mall/pkg/http"
)

func main() {
	s, f, err := mallcmd.NewApp()

	if err != nil {
		panic(err)
	}

	http.Run(s.ServerHTTP, ":3000")

	defer f()
}
