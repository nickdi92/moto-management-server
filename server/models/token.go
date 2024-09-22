package models

import "time"

type TokenRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type TokenResponse struct {
	Token    string     `json:"token"`
	ExpireAt *time.Time `json:"expire_at"`
}
