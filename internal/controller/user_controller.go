package controller

import (
	"github.com/3lur/go-mall/internal/service"
)

type userController struct {
	userSrv service.UserService
}

type UserController interface {
}

func NewUserController(userSrv service.UserService) UserController {
	return &userController{userSrv}
}
