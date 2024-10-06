package models

import (
	"time"
)

// ----------------------------------------------------------------------------
// USERS Structs
// ----------------------------------------------------------------------------

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

type UserLoginResponse struct {
	StatusCode int  `json:"status_code"`
	User       User `json:"user"`
}

type GetUserResponse struct {
	StatusCode int  `json:"status_code"`
	User       User `json:"user"`
}

type UpdateUserResponse struct {
	StatusCode int  `json:"status_code"`
	User       User `json:"user"`
}

// ----------------------------------------------------------------------------
// MOTORCYCLES Structs
// ----------------------------------------------------------------------------

type GetMotorcycleByLicensePlateResponse struct {
	StatusCode int        `json:"status_code"`
	Motorcycle Motorcycle `json:"motorcycle"`
}

type DeleteMotorcycleResponse struct {
	StatusCode int  `json:"status_code"`
	IsDeleted  bool `json:"is_deleted"`
}

type AddMotorcycleResponse struct {
	StatusCode  int          `json:"status_code"`
	Motorcycles []Motorcycle `json:"motorcycles" validate:"required"`
}

type MotorcyclesIndexResponse struct {
	StatusCode  int          `json:"status_code"`
	Motorcycles []Motorcycle `json:"motorcycles" validate:"required"`
}

// ----------------------------------------------------------------------------
// SERVICES Structs
// ----------------------------------------------------------------------------

type AddServiceToMotorcycleResponse struct {
	StatusCode int  `json:"status_code"`
	User       User `json:"user"`
}

type DeleteServiceResponse struct {
	StatusCode int `json:"status_code"`
}

type UpdateServiceToMotorcycleResponse struct {
	StatusCode int  `json:"status_code"`
	User       User `json:"user"`
}

// ----------------------------------------------------------------------------
// FUELS Structs
// ----------------------------------------------------------------------------

type AddFuelToMotorcycleResponse struct {
	StatusCode int  `json:"status_code"`
	User       User `json:"user"`
}
