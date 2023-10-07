package response

import (
	"errors"
	"net/http"

	"github.com/3lur/go-mall/internal/common/reason"
	"github.com/3lur/go-mall/pkg/console"
	"github.com/3lur/go-mall/pkg/e"
	"github.com/gin-gonic/gin"
)

type ResponseBody struct {
	Code    int    `json:"code"`
	Reason  string `json:"reason"`
	Message string `json:"msg"`
	Data    any    `json:"data"`
}

func Build(ctx *gin.Context, err error, data any) {
	if err == nil {
		ctx.JSON(http.StatusOK, buildSuccessResponse(reason.Success, data))
		return
	}

	var myErr *e.Error

	// 未知错误
	if !errors.As(err, &myErr) {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, buildFailResponse(nil))
		return
	}

	// 内部错误
	if e.IsInternalServer(myErr) {
		console.ErrorIf(myErr)
	}

	body := buildFailResponse(myErr)

	if data != nil {
		body.Data = data
	}

	ctx.AbortWithStatusJSON(myErr.Code, body)
}

// buildSuccessResponse
func buildSuccessResponse(reason string, data any) *ResponseBody {
	return &ResponseBody{
		Code:    http.StatusOK,
		Reason:  reason,
		Message: "success",
		Data:    data,
	}
}

// buildFailResponse
func buildFailResponse(err *e.Error) *ResponseBody {
	if err == nil {
		return &ResponseBody{
			Code:    http.StatusInternalServerError,
			Reason:  reason.UnknownError,
			Message: "unknown error",
			Data:    nil,
		}
	}
	return &ResponseBody{
		Code:    err.Code,
		Reason:  err.Reason,
		Message: err.Message,
		Data:    nil,
	}
}
