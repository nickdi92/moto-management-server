package models

import "time"

type TokenRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type DeleteMotorcycleRequest struct {
	Username     string `json:"username" validate:"required"`
	LicensePlate string `json:"license_plate" validate:"required"`
}

type AddFuelToMotorcycleRequest struct {
	Username     string       `json:"username" validate:"required"`
	LicensePlate string       `json:"license_plate" validate:"required"`
	FuelSupplies FuelSupplies `json:"fuel_supplies" validate:"required"`
}

type GetMotorcycleByLicensePlateRequest struct {
	Username     string `json:"username" validate:"required"`
	LicensePlate string `json:"license_plate" validate:"required"`
}

type CreateUserRequest struct {
	Username   string     `json:"username" validate:"required"`
	Password   string     `json:"password" validate:"required"`
	Email      string     `json:"email" validate:"email,required"`
	Name       string     `json:"name" validate:"required"`
	Lastname   string     `json:"lastname" validate:"required"`
	Token      string     `json:"token"`
	ExpireAt   *time.Time `json:"expire_at"`
	IsLoggedIn bool       `json:"is_logged_in"`
}

type AddMotorcycleRequest struct {
	Username    string       `json:"username" validate:"required"`
	Motorcycles []Motorcycle `json:"motorcycles" validate:"required"`
}

type UserLoginRequest struct {
	Username   string     `json:"username" validate:"required"`
	Password   string     `json:"password" validate:"required"`
	Token      string     `json:"token"`
	ExpireAt   *time.Time `json:"expire_at"`
	IsLoggedIn bool       `json:"is_logged_in"`
}

type GetUserRequest struct {
	Username string `json:"username" validate:"required"`
}
