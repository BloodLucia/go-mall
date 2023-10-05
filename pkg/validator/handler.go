package validator

import (
	"github.com/3lur/go-mall/pkg/console"
	"github.com/gin-gonic/gin"
)

func ShouldBindJSON(ctx *gin.Context, body any) bool {
	if err := ctx.ShouldBindJSON(body); err != nil {
		console.Warning("")
		return false
	}

	return true
}
