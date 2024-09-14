package database

import (
	"time"
)

type Motorcycle struct {
	Name         string          `bson:"name"`
	Model        MotorcycleModel `bson:"model"`
	LicensePlate string          `bson:"license_plate"` // Targa
	IsNew        bool            `bson:"is_new"`
	BoughtAt     time.Time       `bson:"bought_at"`
	UpdatedAt    time.Time       `bson:"updated_at"`
}

type MotorcycleModel struct {
	Year         string  `bson:"year"`
	Brand        string  `bson:"brand"`
	Cycling      string  `bson:"cycling"`       //cilindrata
	TankCapacity float32 `bson:"tank_capacity"` // Capienza serbatoio
}
