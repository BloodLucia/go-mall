package schema

import "time"

type UserRegisterRequest struct {
	Username string `json:"user_name" validate:"required,gt=3,lte=30"`
	Email    string `json:"email" validate:"required,email,gt=0,lte=500"`
	Passwd   string `json:"passwd" validate:"required,gte=8,lte=32"`
}

type UserLoginResponse struct {
	ID        string    `json:"id"`
	Username  string    `json:"user_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
