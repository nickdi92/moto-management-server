package models

import (
	"time"
)

// Structs
type Motorcycle struct {
	ID                  string              `json:"id"`
	LicensePlate        string              `json:"license_plate" validate:"required"` // Targa
	MotorcycleDataSheet MotorcycleDataSheet `json:"motorcycle_data_sheet" validate:"required"`
	FuelSupplies        []FuelSupplies      `json:"fuel_supplies"`   // @TODO Gestione rifornimenti
	Service             []Service           `json:"service"`         // @TODO Gestione tagliandi
	Inspection          []Inspection        `json:"inspection"`      // @TODO Gestione Revisioni
	AccidentReport      []AccidentReport    `json:"accident_report"` // @TODO Gestione Incidenti
	CreatedAt           *time.Time          `json:"created_at"`
	UpdatedAt           *time.Time          `json:"updated_at"`
}

type MotorcycleDataSheet struct {
	Name               string    `json:"brand_name" validate:"required"`
	Model              string    `json:"brand_model" validate:"required"`
	ModelYear          string    `json:"model_year" validate:"required"`
	EngineDisplacement string    `json:"engine_displacement" validate:"required"` // Cilindrata
	TankCapacity       string    `json:"tank_capacity" validate:"required"`       // Capacità serbatoio
	Kilometers         string    `json:"kilometers" validate:"required"`
	Insurance          Insurance `json:"insurance"` // Assicurazione
}

type FuelSupplies struct {
	ID            string        `json:"id" validate:"omitempty"`
	PetrolStation PetrolStation `json:"petrol_station"`
	FullFuel      bool          `json:"full_fuel"`
	CreatedAt     string        `json:"created_at"`
}

type Insurance struct {
	IsActive   bool    `json:"is_active"`
	Company    string  `json:"company"`
	PriceMoney float64 `json:"price_money"`
	Currency   string  `json:"currency"`
	Details    string  `json:"details"`
	ExpireAt   string  `json:"expire_at"`
}

type Service struct {
	ID              string        `json:"id"`
	Name            string        `json:"name"`
	LocationAddress Address       `json:"address"`
	ListOfDones     []ListOfDones `json:"list_of_dones"`
	VatPrice        float64       `json:"vat_price"`
	TotalPrice      float64       `json:"total_price"`
	Kilometers      string        `json:"kilometers"`
	ManpowerPrice   float64       `json:"manpower_price"`
	ManpowerHours   int32         `json:"manpower_hours"`
	Date            string        `json:"date"`
}

type ListOfDones struct {
	Name  string  `json:"name"`
	Note  string  `json:"note"`
	Price float64 `json:"price"`
}

type Inspection struct{}

type AccidentReport struct{}

type PetrolStation struct {
	Name               string  `json:"name"`
	Street             string  `json:"street"`
	City               string  `json:"city"`
	Province           string  `json:"province"`
	State              string  `json:"state"`
	FuelType           string  `json:"fuel_type"`
	FuelPricePerLitres float64 `json:"fuel_price_per_litres"`
	TotalLitres        float64 `json:"total_litres"`
	TotalPrice         float64 `json:"total_price"`
}
