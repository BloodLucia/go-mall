package service

import (
	"context"

	"github.com/3lur/go-mall/internal/repo"
	"github.com/3lur/go-mall/internal/schema"
)

type UserService interface {
	RegisterByEmail(ctx context.Context, schema *schema.UserRegisterRequest) error
}

type userService struct {
	userRepo repo.UserRepo
}

// RegisterByEmail implements UserService.
func (*userService) RegisterByEmail(ctx context.Context, schema *schema.UserRegisterRequest) error {
	panic("unimplemented")
}

func NewUserService(userRepo repo.UserRepo) UserService {
	return &userService{userRepo}
}
