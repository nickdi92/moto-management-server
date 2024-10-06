package models

import "time"

// ----------------------------------------------------------------------------
// USERS Structs
// ----------------------------------------------------------------------------

type CreateUserRequest struct {
	Username     string     `json:"username" validate:"required"`
	Password     string     `json:"password" validate:"required"`
	Email        string     `json:"email" validate:"email,required"`
	Name         string     `json:"name" validate:"required"`
	Lastname     string     `json:"lastname" validate:"required"`
	Token        string     `json:"token"`
	ExpireAt     *time.Time `json:"expire_at"`
	IsLoggedIn   bool       `json:"is_logged_in"`
	Address      Address    `json:"address"`
	UserRegistry Registry   `json:"registry"`
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

type UpdateUserRequest struct {
	CreateUserRequest CreateUserRequest `json:"update_user"`
}

type TokenRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// ----------------------------------------------------------------------------
// MOTORCYCLES Structs
// ----------------------------------------------------------------------------

type DeleteMotorcycleRequest struct {
	Username     string `json:"username" validate:"required"`
	LicensePlate string `json:"license_plate" validate:"required"`
}

type AddMotorcycleRequest struct {
	Username    string       `json:"username" validate:"required"`
	Motorcycles []Motorcycle `json:"motorcycles" validate:"required"`
}

type GetMotorcycleByLicensePlateRequest struct {
	Username     string `json:"username" validate:"required"`
	LicensePlate string `json:"license_plate" validate:"required"`
}

type MotorcyclesIndexRequest struct {
	Username string `json:"username" validate:"required"`
}

// ----------------------------------------------------------------------------
// FUEL Structs
// ----------------------------------------------------------------------------

type AddFuelToMotorcycleRequest struct {
	Username     string       `json:"username" validate:"required"`
	LicensePlate string       `json:"license_plate" validate:"required"`
	FuelSupplies FuelSupplies `json:"fuel_supplies" validate:"required"`
}

// ----------------------------------------------------------------------------
// SERVICE Structs
// ----------------------------------------------------------------------------

type AddServiceToMotorcycleRequest struct {
	Username     string  `json:"username" validate:"required"`
	LicensePlate string  `json:"license_plate" validate:"required"`
	Service      Service `json:"service" validate:"required"`
}

type DeleteServiceRequest struct {
	Username     string `json:"username" validate:"required"`
	LicensePlate string `json:"license_plate" validate:"required"`
	ServiceId    string `json:"service_id" validate:"required"`
}

type UpdateServiceRequest struct {
	Username     string  `json:"username" validate:"required"`
	LicensePlate string  `json:"license_plate" validate:"required"`
	Service      Service `json:"service" validate:"required"`
}
