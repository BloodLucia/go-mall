package server

import "github.com/gin-gonic/gin"

func NewServerHTTP() *gin.Engine {
	r := gin.New()

	return r

}
