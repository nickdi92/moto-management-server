package models

import (
	"time"

	"github.com/govalues/money"
)

type FuelType string

const FuelType_Gasoline = "gasoline"
const FuelType_Diesel = "diesel"

type Motorcycle struct {
	ID                  string              `json:"id"`
	LicensePlate        string              `json:"license_plate"` // Targa
	MotorcycleDataSheet MotorcycleDataSheet `json:"motorcycle_data_sheet"`
	FuelSupplies        []FuelSupplies      `json:"fuel_supplies"`   // @TODO Gestione rifornimenti
	Service             []Service           `json:"service"`         // @TODO Gestione tagliandi
	Inspection          []Inspection        `json:"inspection"`      // @TODO Gestione Revisioni
	AccidentReport      []AccidentReport    `json:"accident_report"` // @TODO Gestione Incidenti
	CreatedAt           *time.Time          `json:"created_at"`
	UpdatedAt           *time.Time          `json:"updated_at"`
}

type MotorcycleDataSheet struct {
	Name               string    `json:"brand_name"`
	Model              string    `json:"brand_model"`
	ModelYear          string    `json:"model_year"`
	EngineDisplacement string    `json:"engine_displacement"` // Cilindrata
	TankCapacity       string    `json:"tank_capacity"`       // Capacità serbatoio
	Kilometers         string    `json:"kilometers"`
	Insurance          Insurance `json:"insurance"` // Assicurazione
}

type FuelSupplies struct {
	PetrolStation PetrolStation `json:"location"`
	FullFuel      bool          `json:"full_fuel"`
	CreatedAt     *time.Time    `json:"json"`
}

type PetrolStation struct {
	Name               string       `json:"name"`
	Street             string       `json:"street"`
	City               string       `json:"city"`
	Province           string       `json:"province"`
	State              string       `json:"state"`
	FuelType           FuelType     `json:"fuel_type"`
	FuelPricePerLitres money.Amount `json:"fuel_price_per_litres"`
	TotalLitres        float64      `json:"total_litres"`
	TotalPrice         money.Amount `json:"total_price"`
}

type Insurance struct {
	IsActive   bool         `json:"is_active"`
	Company    string       `json:"company"`
	PriceMoney money.Amount `json:"price_money"`
	Currency   string       `json:"currency"`
	Details    string       `json:"details"`
	ExpireAt   *time.Time   `json:"expire_at"`
}

type Service struct{}

type Inspection struct{}

type AccidentReport struct{}
