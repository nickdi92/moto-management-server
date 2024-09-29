package models

import (
	"time"
)

type User struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Lastname     string       `json:"lastname"`
	Username     string       `json:"username"`
	Password     string       `json:"password"`
	Email        string       `json:"email"`
	Token        string       `json:"token"`
	ExpireAt     *time.Time   `json:"expire_at"`
	CreatedAt    *time.Time   `json:"created_at"`
	UpdatedAt    *time.Time   `json:"updated_at"`
	IsLoggedIn   bool         `json:"is_logged_in"`
	Motorcycles  []Motorcycle `json:"motorcycles"`
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

func (u User) MergeMotorcyclesIDS(oldUser User) {
	if oldUser.Motorcycles == nil || len(oldUser.Motorcycles) == 0 {
		return
	}
	for index, newMt := range u.Motorcycles {
		if oldMt := oldUser.Motorcycles[index]; oldMt.LicensePlate == newMt.LicensePlate {
			u.Motorcycles[index].ID = oldMt.ID
		}
	}
}
