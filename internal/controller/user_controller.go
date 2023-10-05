package controller

type userController struct {
}

type UserController interface {
}

func NewUserController() UserController {
	return &userController{}
}
