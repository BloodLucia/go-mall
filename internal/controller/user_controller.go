package controller

import (
	"github.com/3lur/go-mall/internal/common/reason"
	"github.com/3lur/go-mall/internal/schema"
	"github.com/3lur/go-mall/internal/service"
	"github.com/3lur/go-mall/pkg/e"
	"github.com/3lur/go-mall/pkg/response"
	"github.com/gin-gonic/gin"
)

type userController struct {
	userSrv service.UserService
}

// Register implements UserController.
func (uc *userController) Register(ctx *gin.Context) {
	var req schema.UserRegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Build(ctx, e.BadRequest(reason.RequestBodyError).WithErr(err), nil)
		return
	}
	resp := uc.userSrv.RegisterByEmail(ctx, &req)
	response.Build(ctx, nil, resp)
}

type UserController interface {
	Register(ctx *gin.Context)
}

func NewUserController(userSrv service.UserService) UserController {
	return &userController{userSrv}
}
