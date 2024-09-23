package models

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

type GetUserRoute struct {
	Username string `json:"username" validate:"required"`
}

type User struct {
	ID          string       `json:"id"`
	Username    string       `json:"username"`
	Password    string       `json:"password"`
	Email       string       `json:"email"`
	Name        string       `json:"name"`
	Lastname    string       `json:"lastname"`
	Motorcycles []Motorcycle `json:"motorcycles"`
	Token       string       `json:"token"`
	ExpireAt    *time.Time   `json:"expire_at"`
	IsLoggedIn  bool         `json:"is_logged_in"`
}
