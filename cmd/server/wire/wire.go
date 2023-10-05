//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/3lur/go-mall/internal/server"
	"github.com/3lur/go-mall/pkg/config"
	"github.com/google/wire"
)

func NewApp(*config.Config) (*server.Server, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
	))
}
