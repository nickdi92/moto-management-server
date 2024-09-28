package models

import (
	"moto-management-server/database/models"
)

func (ft FuelType) ToString() string {
	return string(ft)
}

func (f FuelSupplies) ToMongoFuelSupplies() models.FuelSupplies {
	floatFuelPricePerLitres, _ := f.PetrolStation.FuelPricePerLitres.Float64()
	floatFuelTotalPrice, _ := f.PetrolStation.TotalPrice.Float64()

	return models.FuelSupplies{
		PetrolStation: models.PetrolStation{
			Name:               f.PetrolStation.Name,
			Street:             f.PetrolStation.Street,
			City:               f.PetrolStation.City,
			Province:           f.PetrolStation.Province,
			State:              f.PetrolStation.State,
			FuelType:           f.PetrolStation.FuelType.ToString(),
			FuelPricePerLitres: floatFuelPricePerLitres,
			TotalLitres:        f.PetrolStation.TotalLitres,
			TotalPrice:         floatFuelTotalPrice,
		},
		FullFuel:  f.FullFuel,
		CreatedAt: f.CreatedAt,
	}
}
