//go:build wireinject
// +build wireinject

package main

import (
	"github.com/3lur/go-mall/internal/server"
	"github.com/google/wire"
)

func newApp() (*server.Server, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
	))
}
