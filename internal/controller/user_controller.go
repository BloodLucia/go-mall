package controller

import "github.com/3lur/go-mall/internal/common/data"

type userController struct {
	*data.Data
}

type UserController interface {
}

func NewUserController(data *data.Data) UserController {
	return &userController{data}
}
