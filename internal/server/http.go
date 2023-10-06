package server

import (
	"github.com/3lur/go-mall/internal/controller"
	"github.com/gin-gonic/gin"
)

func NewServerHTTP(
	userC controller.UserController,
	pingC controller.PingController,
) *gin.Engine {
	r := gin.New()

	r.GET("/ping", pingC.Ping)

	r.POST("/register", userC.Register)

	return r

}
