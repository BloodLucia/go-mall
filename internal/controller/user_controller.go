package controller

import (
	"github.com/3lur/go-mall/internal/repo"
)

type userController struct {
	userRepo repo.UserRepo
}

type UserController interface {
}

func NewUserController(userRepo repo.UserRepo) UserController {
	return &userController{userRepo}
}
