package controller

import (
	"github.com/3lur/go-mall/internal/schema"
	"github.com/3lur/go-mall/internal/service"
	"github.com/3lur/go-mall/pkg/response"
	"github.com/3lur/go-mall/pkg/validator"
	"github.com/gin-gonic/gin"
)

type userController struct {
	userSrv service.UserService
}

// Register implements UserController.
func (uc *userController) Register(ctx *gin.Context) {
	var req schema.UserRegisterRequest
	if !validator.BindAndCheck(ctx, &req) {
		return
	}
	response.Build(ctx, nil, "ok!")
}

type UserController interface {
	Register(ctx *gin.Context)
}

func NewUserController(userSrv service.UserService) UserController {
	return &userController{userSrv}
}
