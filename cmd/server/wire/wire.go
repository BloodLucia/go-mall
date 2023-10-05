//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/3lur/go-mall/internal/common/data"
	"github.com/3lur/go-mall/internal/controller"
	"github.com/3lur/go-mall/internal/repo"
	"github.com/3lur/go-mall/internal/server"
	"github.com/3lur/go-mall/pkg/config"
	"github.com/google/wire"
)

func NewApp(conf *config.Config) (*server.Server, func(), error) {
	panic(wire.Build(
		data.NewData,
		repo.ProviderSet,
		controller.ProviderSet,
		server.ProviderSet,
	))
}
