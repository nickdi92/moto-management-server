package models

import "time"

type User struct {
	ID           string       `json:"id"`
	Username     string       `json:"username"`
	Password     string       `json:"password"`
	Email        string       `json:"email"`
	Name         string       `json:"name"`
	Lastname     string       `json:"lastname"`
	Motorcycles  []Motorcycle `json:"motorcycles"`
	Token        string       `json:"token"`
	ExpireAt     *time.Time   `json:"expire_at"`
	IsLoggedIn   bool         `json:"is_logged_in"`
	Address      Address      `json:"address"`
	UserRegistry Registry     `json:"registry"`
}

type Address struct {
	City     string `json:"city"`
	Street   string `json:"street"`
	ZipCode  string `json:"zip_code"`
	Province string `json:"province"`
	State    string `json:"state"`
}

// Anagrafrica
type Registry struct {
	FiscalCode string `json:"fiscal_code"`
	DOB        string `json:"dob"`
}
