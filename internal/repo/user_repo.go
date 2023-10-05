package repo

import (
	"context"
	"errors"

	"github.com/3lur/go-mall/internal/common/data"
	"github.com/3lur/go-mall/internal/model"
)

type UserRepo interface {
	FindByEmail(ctx context.Context, email string) (user *model.User, exist bool, err error)
	FindByUsername(ctx context.Context, username string) (user *model.User, exist bool, err error)
}

type userRepo struct {
	data *data.Data
}

// FindByUsername get user by username.
func (ur *userRepo) FindByUsername(ctx context.Context, username string) (user *model.User, exist bool, err error) {
	user = &model.User{}
	exist, err = ur.data.DB.Context(ctx).Where("username = ?", username).Get(user)
	if err != nil {
		err = errors.New("用户不存在")
	}
	return
}

// FindByEmail get user by email.
func (ur *userRepo) FindByEmail(ctx context.Context, email string) (user *model.User, exist bool, err error) {
	user = &model.User{}
	exist, err = ur.data.DB.Context(ctx).Where("email = ?", email).Get(user)
	if err != nil {
		err = errors.New("用户不存在")
	}
	return
}

func NewUserRepo(data *data.Data) UserRepo {
	return &userRepo{data}
}
