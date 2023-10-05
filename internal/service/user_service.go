package service

import (
	"context"

	"github.com/3lur/go-mall/internal/repo"
	"github.com/3lur/go-mall/internal/schema"
)

type UserService interface {
	RegisterByEmail(ctx context.Context, schema *schema.UserRegisterRequest) (resp *schema.UserLoginResponse)
}

type userService struct {
	userRepo repo.UserRepo
}

// RegisterByEmail register user with email.
func (us *userService) RegisterByEmail(ctx context.Context, schema *schema.UserRegisterRequest) (resp *schema.UserLoginResponse) {
	_, has, err := us.userRepo.FindByEmail(ctx, schema.Email)
	if err != nil {
		return nil
	}
	if has {
		return nil
	}
	return nil
}

func NewUserService(userRepo repo.UserRepo) UserService {
	return &userService{userRepo}
}
