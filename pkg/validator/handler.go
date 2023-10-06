package validator

import (
	"fmt"
	"net/http"

	"github.com/3lur/go-mall/internal/common/reason"
	"github.com/3lur/go-mall/pkg/console"
	"github.com/3lur/go-mall/pkg/e"
	"github.com/3lur/go-mall/pkg/response"
	"github.com/gin-gonic/gin"
)

func BindAndCheck(ctx *gin.Context, data any) bool {
	// 支持 JSON、Form、URL Query
	if err := ctx.ShouldBind(data); err != nil {
		console.Warning(fmt.Sprintf("http_handle ShouldBind error: %s", err))
		response.Build(ctx, e.New(http.StatusBadRequest, reason.RequestBodyError), nil)
		return false
	}

	errFields, err := Validate.Check(data)

	if err != nil {
		response.Build(ctx, err, errFields)
		return false
	}

	return true
}
