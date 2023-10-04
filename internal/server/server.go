package server

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type Server struct {
	ServerHTTP *gin.Engine
}

func NewServer(ServerHTTP *gin.Engine) *Server {
	return &Server{ServerHTTP}
}

var ProviderSet = wire.NewSet(
	NewServer,
	NewServerHTTP,
)
