package main

import (
	"github.com/3lur/go-mall/pkg/http"
)

func main() {
	s, f, err := newApp()
	if err != nil {
		panic(err)
	}

	http.Run(s.ServerHTTP, ":3000")

	defer f()
}
