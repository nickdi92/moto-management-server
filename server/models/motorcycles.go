package models

import (
	"golang.org/x/text/currency"
	"time"
)

type Motorcycle struct {
	ID                  string              `json:"id"`
	LicensePlate        string              `json:"license_plate" validate:"required"` // Targa
	MotorcycleDataSheet MotorcycleDataSheet `json:"motorcycle_data_sheet" validate:"required"`
	FuelSupplies        FuelSupplies        `json:"fuel_supplies"`   // @TODO Gestione rifornimenti
	Service             Service             `json:"service"`         // @TODO Gestione tagliandi
	Inspection          Inspection          `json:"inspection"`      // @TODO Gestione Revisioni
	AccidentReport      AccidentReport      `json:"accident_report"` // @TODO Gestione Incidenti
	CreatedAt           *time.Time          `json:"created_at"`
	UpdatedAt           *time.Time          `json:"updated_at"`
}

type MotorcycleDataSheet struct {
	Name               string    `json:"brand_name" validate:"required"`
	Model              string    `json:"brand_model" validate:"required"`
	ModelYear          string    `json:"model_year" validate:"required"`
	EngineDisplacement string    `json:"engine_displacement" validate:"required"` // Cilindrata
	TankCapacity       string    `json:"tank_capacity" validate:"required"`       // Capacità serbatoio
	Insurance          Insurance `json:"insurance"`                               // Assicurazione
}

type FuelSupplies struct{}

type Insurance struct {
	IsActive   bool            `json:"is_active"`
	Company    string          `json:"company"`
	PriceMoney currency.Amount `json:"price_money"`
	Details    string          `json:"details"`
	ExpireAt   *time.Time      `json:"expire_at"`
}

type Service struct{}

type Inspection struct{}

type AccidentReport struct{}
