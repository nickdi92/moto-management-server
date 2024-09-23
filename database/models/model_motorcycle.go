package models

import (
	"golang.org/x/text/currency"
	"time"
)

type Motorcycle struct {
	ID                  string              `json:"id"`
	LicensePlate        string              `json:"license_plate" bson:"license_plate"` // Targa
	MotorcycleDataSheet MotorcycleDataSheet `json:"motorcycle_data_sheet" bson:"motorcycle_data_sheet"`
	FuelSupplies        FuelSupplies        `json:"fuel_supplies" bson:"fuel_supplies"`     // @TODO Gestione rifornimenti
	Service             Service             `json:"service" bson:"service"`                 // @TODO Gestione tagliandi
	Inspection          Inspection          `json:"inspection" bson:"inspection"`           // @TODO Gestione Revisioni
	AccidentReport      AccidentReport      `json:"accident_report" bson:"accident_report"` // @TODO Gestione Incidenti
	CreatedAt           *time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt           *time.Time          `json:"updated_at" bson:"updated_at"`
}

type MotorcycleDataSheet struct {
	Name               string    `json:"brand_name" bson:"name"`
	Model              string    `json:"brand_model" bson:"model"`
	ModelYear          string    `json:"model_year" bson:"model_year"`
	EngineDisplacement string    `json:"engine_displacement" bson:"engine_displacement"` // Cilindrata
	TankCapacity       string    `json:"tank_capacity" bson:"tank_capacity"`             // Capacità serbatoio
	Insurance          Insurance `json:"insurance" bson:"insurance"`                     // Assicurazione
}

type FuelSupplies struct{}

type Insurance struct {
	IsActive   bool            `json:"is_active" bson:"is_active"`
	Company    string          `json:"company" bson:"company"`
	PriceMoney currency.Amount `json:"price_money" bson:"price_money"`
	Details    string          `json:"details" bson:"details"`
	ExpireAt   *time.Time      `json:"expire_at" bson:"expire_at"`
}

type Service struct{}

type Inspection struct{}

type AccidentReport struct{}
