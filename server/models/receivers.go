package models

import (
	"moto-management-server/business_logic/models"
	models2 "moto-management-server/business_logic/models"
	"time"

	"github.com/govalues/money"
)

func (f FuelSupplies) ToBusinessLogicModel() models2.FuelSupplies {
	pricePerLitres, _ := money.NewAmountFromFloat64(money.EUR.Code(), f.PetrolStation.FuelPricePerLitres)
	totalPrice, _ := money.NewAmountFromFloat64(money.EUR.Code(), f.PetrolStation.TotalPrice)
	createdAt, _ := time.Parse(time.DateOnly, f.CreatedAt)

	model := models2.FuelSupplies{
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
	return model
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

func (s Service) ToBusinessLogicModel() models2.Service {
	listOfDones := make([]models2.ListOfDones, 0)
	for _, stuff := range s.ListOfDones {
		stuffPrice, _ := money.NewAmountFromFloat64(money.EUR.Code(), stuff.Price)
		listOfDones = append(listOfDones, models2.ListOfDones{
			ID:    stuff.ID,
			Name:  stuff.Name,
			Note:  stuff.Note,
			Price: stuffPrice,
		})
	}
	vatPrice, _ := money.NewAmountFromFloat64(money.EUR.Code(), s.VatPrice)
	totalPrice, _ := money.NewAmountFromFloat64(money.EUR.Code(), s.TotalPrice)
	manpowerPrice, _ := money.NewAmountFromFloat64(money.EUR.Code(), s.ManpowerPrice)
	date, _ := time.Parse(time.DateOnly, s.Date)
	return models2.Service{
		ID:   s.ID,
		Name: s.Name,
		LocationAddress: models2.Address{
			City:     s.LocationAddress.City,
			Street:   s.LocationAddress.Street,
			ZipCode:  s.LocationAddress.ZipCode,
			Province: s.LocationAddress.Province,
			State:    s.LocationAddress.State,
		},
		Kilometers:    s.Kilometers,
		ListOfDones:   listOfDones,
		VatPrice:      vatPrice,
		TotalPrice:    totalPrice,
		ManpowerPrice: manpowerPrice,
		ManpowerHours: s.ManpowerHours,
		Date:          date,
	}
}

func (s Service) ToServerModel(bs models.Service) Service {
	listOfDones := make([]ListOfDones, 0)
	for _, stuff := range bs.ListOfDones {
		stuffPrice, _ := stuff.Price.Float64()
		listOfDones = append(listOfDones, ListOfDones{
			ID:    stuff.ID,
			Name:  stuff.Name,
			Note:  stuff.Note,
			Price: stuffPrice,
		})
	}
	vatPrice, _ := bs.VatPrice.Float64()
	totalPrice, _ := bs.TotalPrice.Float64()
	manpowerPrice, _ := bs.ManpowerPrice.Float64()
	return Service{
		ID:   bs.ID,
		Name: bs.Name,
		LocationAddress: Address{
			City:     bs.LocationAddress.City,
			Street:   bs.LocationAddress.Street,
			ZipCode:  bs.LocationAddress.ZipCode,
			Province: bs.LocationAddress.Province,
			State:    bs.LocationAddress.State,
		},
		Kilometers:    bs.Kilometers,
		ListOfDones:   listOfDones,
		VatPrice:      vatPrice,
		TotalPrice:    totalPrice,
		ManpowerPrice: manpowerPrice,
		ManpowerHours: s.ManpowerHours,
		Date:          bs.Date.String(),
	}
}
