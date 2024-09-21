package server

import "time"

type RegisterUserRequest struct {
	Username   string     `json:"username" validate:"required"`
	Password   string     `json:"password" validate:"required"`
	Email      string     `json:"email" validate:"email,required"`
	Name       string     `json:"name" validate:"required"`
	Lastname   string     `json:"lastname" validate:"required"`
	Token      string     `json:"token"`
	ExpireAt   *time.Time `json:"expire_at"`
	IsLoggedIn bool       `json:"is_logged_in"`
}

type UserLoginRequest struct {
	Username   string     `json:"username" validate:"required"`
	Password   string     `json:"password" validate:"required"`
	Token      string     `json:"token"`
	ExpireAt   *time.Time `json:"expire_at"`
	IsLoggedIn bool       `json:"is_logged_in"`
}

type MotorcycleCreateRequest struct {
}

type TokenRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type TokenResponse struct {
	Token    string     `json:"token"`
	ExpireAt *time.Time `json:"expire_at"`
}
