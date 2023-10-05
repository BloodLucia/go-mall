package service

import "github.com/3lur/go-mall/internal/repo"

type UserService interface {
}

type userService struct {
	userRepo repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) UserService {
	return &userService{userRepo}
}
