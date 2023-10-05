package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingController interface {
	Ping(ctx *gin.Context)
}

type pingController struct {
}

func (*pingController) Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong!")
}

func NewPingController() PingController {
	return &pingController{}
}
