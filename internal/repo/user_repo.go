package repo

import "github.com/3lur/go-mall/internal/common/data"

type UserRepo interface {
}

type userRepo struct {
	data *data.Data
}

func NewUserRepo(data *data.Data) UserRepo {
	return &userRepo{data}
}
