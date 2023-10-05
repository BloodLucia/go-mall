//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/3lur/go-mall/internal/server"
	"github.com/google/wire"
)

func NewApp() (*server.Server, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
	))
}
