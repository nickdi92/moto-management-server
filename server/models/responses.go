package models

import "time"

type DeleteMotorcycleResponse struct {
	StatusCode int  `json:"status_code"`
	IsDeleted  bool `json:"is_deleted"`
}

type AddFuelToMotorcycleResponse struct {
	StatusCode int        `json:"status_code"`
	Motorcycle Motorcycle `json:"motorcycle"`
}

type GetMotorcycleByLicensePlateResponse struct {
	StatusCode int        `json:"status_code"`
	Motorcycle Motorcycle `json:"motorcycle"`
}

type TokenResponse struct {
	StatusCode int        `json:"status_code"`
	Token      string     `json:"token"`
	ExpireAt   *time.Time `json:"expire_at"`
}

type CreateUserResponse struct {
	StatusCode int        `json:"status_code"`
	Token      string     `json:"token"`
	ExpireAt   *time.Time `json:"expire_at"`
}

type AddMotorcycleResponse struct {
	StatusCode  int          `json:"status_code"`
	Motorcycles []Motorcycle `json:"motorcycles" validate:"required"`
}

type UserLoginResponse struct {
	StatusCode int        `json:"status_code"`
	Token      string     `json:"token"`
	ExpireAt   *time.Time `json:"expire_at"`
	IsLoggedIn bool       `json:"is_logged_in"`
}

type GetUserResponse struct {
	StatusCode int  `json:"status_code"`
	User       User `json:"user"`
}
