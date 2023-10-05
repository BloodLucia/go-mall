package server

import (
	"github.com/3lur/go-mall/internal/controller"
	"github.com/gin-gonic/gin"
)

func NewServerHTTP(
	userC controller.UserController,
) *gin.Engine {
	r := gin.New()

	return r

}
