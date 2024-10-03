package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"moto-management-server/database/models"
)

func (ft FuelType) ToString() string {
	return string(ft)
}

func (f FuelSupplies) ToMongoFuelSupplies() models.FuelSupplies {
	floatFuelPricePerLitres, _ := f.PetrolStation.FuelPricePerLitres.Float64()
	floatFuelTotalPrice, _ := f.PetrolStation.TotalPrice.Float64()
	id, _ := primitive.ObjectIDFromHex(f.ID)
	return models.FuelSupplies{
		ID: id,
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

func (s Service) ToMongoService() models.Service {
	listOfDones := make([]models.ListOfDones, 0)
	for _, stuff := range s.ListOfDones {
		stuffPrice, _ := stuff.Price.Float64()
		mongoListOfDone := models.ListOfDones{
			Name:  stuff.Name,
			Note:  stuff.Note,
			Price: stuffPrice,
		}
		if s.ID != "" {
			mongoListOfDone.ID, _ = primitive.ObjectIDFromHex(s.ID)
		}
		listOfDones = append(listOfDones, mongoListOfDone)
	}
	vatPrice, _ := s.VatPrice.Float64()
	totalPrice, _ := s.TotalPrice.Float64()
	manpowerPrice, _ := s.ManpowerPrice.Float64()
	mongoService := models.Service{
		Name: s.Name,
		LocationAddress: models.Address{
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
		Date:          s.Date,
	}

	if s.ID != "" {
		mongoService.ID, _ = primitive.ObjectIDFromHex(s.ID)
	}

	return mongoService
}
