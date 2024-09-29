package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	Lastname     string             `bson:"lastname"`
	Username     string             `bson:"username"`
	Password     string             `bson:"password"`
	Email        string             `bson:"email"`
	Token        string             `bson:"token"`
	ExpireAt     time.Time          `bson:"expire_at"`
	Motorcycles  []Motorcycle       `bson:"motorcycles"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
	IsLoggedIn   bool               `bson:"is_logged_in"`
	Address      Address            `bson:"address"`
	UserRegistry Registry           `bson:"registry"`
}

type Address struct {
	City     string `bson:"city"`
	Street   string `bson:"street"`
	ZipCode  string `bson:"zip_code"`
	Province string `bson:"province"`
	State    string `bson:"state"`
}

// Anagrafrica
type Registry struct {
	FiscalCode string    `bson:"fiscal_code"`
	DOB        time.Time `bson:"dob"`
}
