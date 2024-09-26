package models

import (
	"github.com/Rhymond/go-money"
	"time"
)

type Motorcycle struct {
	ID                  string              `json:"id"`
	LicensePlate        string              `json:"license_plate"` // Targa
	MotorcycleDataSheet MotorcycleDataSheet `json:"motorcycle_data_sheet"`
	FuelSupplies        FuelSupplies        `json:"fuel_supplies"`   // @TODO Gestione rifornimenti
	Service             Service             `json:"service"`         // @TODO Gestione tagliandi
	Inspection          Inspection          `json:"inspection"`      // @TODO Gestione Revisioni
	AccidentReport      AccidentReport      `json:"accident_report"` // @TODO Gestione Incidenti
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

type FuelSupplies struct{
	
}

type Insurance struct {
	IsActive   bool         `json:"is_active"`
	Company    string       `json:"company"`
	PriceMoney *money.Money `json:"price_money"`
	Currency   string       `json:"currency"`
	Details    string       `json:"details"`
	ExpireAt   *time.Time   `json:"expire_at"`
}

type Service struct{}

type Inspection struct{}

type AccidentReport struct{}
