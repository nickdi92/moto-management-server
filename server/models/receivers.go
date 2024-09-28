package models

import (
	models2 "moto-management-server/business_logic/models"
	"time"

	"github.com/govalues/money"
)

func (f FuelSupplies) ToBusinessLogicModel() models2.FuelSupplies {
	pricePerLitres, _ := money.NewAmountFromFloat64(money.EUR.Code(), f.PetrolStation.FuelPricePerLitres)
	totalPrice, _ := money.NewAmountFromFloat64(money.EUR.Code(), f.PetrolStation.TotalPrice)
	createdAt, _ := time.Parse(time.DateOnly, f.CreatedAt)

	modl := models2.FuelSupplies{
		ID: f.ID,
		PetrolStation: models2.PetrolStation{
			Name:               f.PetrolStation.Name,
			Street:             f.PetrolStation.Street,
			City:               f.PetrolStation.City,
			Province:           f.PetrolStation.Province,
			State:              f.PetrolStation.State,
			FuelType:           models2.FuelType(f.PetrolStation.FuelType),
			FuelPricePerLitres: pricePerLitres,
			TotalLitres:        f.PetrolStation.TotalLitres,
			TotalPrice:         totalPrice,
		},
		FullFuel:  f.FullFuel,
		CreatedAt: &createdAt,
	}
	return modl
}

func (f FuelSupplies) ToServerModel(blModel models2.FuelSupplies) FuelSupplies {
	fuelPrice, _ := blModel.PetrolStation.FuelPricePerLitres.Float64()
	totalPrice, _ := blModel.PetrolStation.TotalPrice.Float64()

	return FuelSupplies{
		ID:        blModel.ID,
		FullFuel:  blModel.FullFuel,
		CreatedAt: blModel.CreatedAt.String(),
		PetrolStation: PetrolStation{
			Name:               blModel.PetrolStation.Name,
			Street:             blModel.PetrolStation.Street,
			City:               blModel.PetrolStation.City,
			Province:           blModel.PetrolStation.Province,
			State:              blModel.PetrolStation.State,
			FuelType:           blModel.PetrolStation.FuelType.ToString(),
			FuelPricePerLitres: fuelPrice,
			TotalLitres:        blModel.PetrolStation.TotalLitres,
			TotalPrice:         totalPrice,
		},
	}
}
