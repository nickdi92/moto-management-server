package business_logic

import (
	"moto-management-server/business_logic/models"
	models2 "moto-management-server/database/models"
	"time"

	money2 "github.com/govalues/money"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func fromMongoUserToBlUser(mongoUser models2.User) models.User {
	id := mongoUser.ID.Hex()
	if mongoUser.ID.IsZero() {
		id = ""
	}
	blUser := models.User{
		ID:          id,
		Username:    mongoUser.Username,
		Name:        mongoUser.Name,
		Lastname:    mongoUser.Lastname,
		Password:    mongoUser.Password,
		Email:       mongoUser.Email,
		Token:       mongoUser.Token,
		ExpireAt:    &mongoUser.ExpireAt,
		CreatedAt:   &mongoUser.CreatedAt,
		UpdatedAt:   &mongoUser.UpdatedAt,
		IsLoggedIn:  mongoUser.IsLoggedIn,
		Motorcycles: fromMongoMotorcyclesToBlMotorcycles(mongoUser.Motorcycles),
		Address: models.Address{
			City:     mongoUser.Address.City,
			Street:   mongoUser.Address.Street,
			ZipCode:  mongoUser.Address.ZipCode,
			Province: mongoUser.Address.Province,
			State:    mongoUser.Address.State,
		},
		UserRegistry: models.Registry{
			FiscalCode: mongoUser.UserRegistry.FiscalCode,
			DOB:        mongoUser.UserRegistry.DOB.String(),
		},
	}
	return blUser
}

func fromBlUserToMongoUser(blUser models.User) models2.User {
	dob, _ := time.Parse(time.DateOnly, blUser.UserRegistry.DOB)
	mongoUser := models2.User{
		Username:    blUser.Username,
		Name:        blUser.Name,
		Lastname:    blUser.Lastname,
		Password:    blUser.Password,
		Email:       blUser.Email,
		Token:       blUser.Token,
		IsLoggedIn:  blUser.IsLoggedIn,
		Motorcycles: fromBlMotorcyclesToMongoMotorcycles(blUser.Motorcycles),
		Address: models2.Address{
			City:     blUser.Address.City,
			Street:   blUser.Address.Street,
			ZipCode:  blUser.Address.ZipCode,
			Province: blUser.Address.Province,
			State:    blUser.Address.State,
		},
		UserRegistry: models2.Registry{
			FiscalCode: blUser.UserRegistry.FiscalCode,
			DOB:        dob,
		},
	}

	if blUser.ExpireAt != nil {
		mongoUser.ExpireAt = *blUser.ExpireAt
	}

	if blUser.ID != "" {
		mongoUser.ID, _ = primitive.ObjectIDFromHex(blUser.ID)
	}

	return mongoUser
}

func fromBlMotorcyclesToMongoMotorcycles(blMotorcycles []models.Motorcycle) []models2.Motorcycle {
	if blMotorcycles != nil {
		mongoMotorcycles := make([]models2.Motorcycle, 0)
		for _, mt := range blMotorcycles {
			priceMoney, _ := mt.MotorcycleDataSheet.Insurance.PriceMoney.Float64()
			mongoMt := models2.Motorcycle{
				LicensePlate: mt.LicensePlate,
				MotorcycleDataSheet: models2.MotorcycleDataSheet{
					Name:               mt.MotorcycleDataSheet.Name,
					Model:              mt.MotorcycleDataSheet.Model,
					ModelYear:          mt.MotorcycleDataSheet.ModelYear,
					EngineDisplacement: mt.MotorcycleDataSheet.EngineDisplacement,
					TankCapacity:       mt.MotorcycleDataSheet.TankCapacity,
					Kilometers:         mt.MotorcycleDataSheet.Kilometers,
					Insurance: models2.Insurance{
						IsActive:   mt.MotorcycleDataSheet.Insurance.IsActive,
						Company:    mt.MotorcycleDataSheet.Insurance.Company,
						PriceMoney: priceMoney,
						Currency:   mt.MotorcycleDataSheet.Insurance.PriceMoney.Curr().Code(),
						Details:    mt.MotorcycleDataSheet.Insurance.Details,
						ExpireAt:   mt.MotorcycleDataSheet.Insurance.ExpireAt,
					},
				},
				FuelSupplies:   fromBlFuelsSuppliesToMongoFuelsSupplies(mt.FuelSupplies),
				Service:        fromBlServicesToMongoServices(mt.Service),
				Inspection:     []models2.Inspection{},
				AccidentReport: []models2.AccidentReport{},
				CreatedAt:      mt.CreatedAt,
				UpdatedAt:      mt.UpdatedAt,
			}

			if mt.ID != "" {
				mongoId, _ := primitive.ObjectIDFromHex(mt.ID)
				mongoMt.ID = mongoId
			}
			mongoMotorcycles = append(mongoMotorcycles, mongoMt)
		}

		return mongoMotorcycles
	}
	return nil
}

func fromMongoMotorcyclesToBlMotorcycles(mongoMotorcycles []models2.Motorcycle) []models.Motorcycle {
	if mongoMotorcycles != nil {
		blMotorcycles := make([]models.Motorcycle, 0)
		for _, mt := range mongoMotorcycles {
			money, _ := money2.NewAmountFromFloat64(mt.MotorcycleDataSheet.Insurance.Currency, mt.MotorcycleDataSheet.Insurance.PriceMoney)
			blMotorcycles = append(blMotorcycles, models.Motorcycle{
				ID:           mt.ID.Hex(),
				LicensePlate: mt.LicensePlate,
				MotorcycleDataSheet: models.MotorcycleDataSheet{
					Name:               mt.MotorcycleDataSheet.Name,
					Model:              mt.MotorcycleDataSheet.Model,
					ModelYear:          mt.MotorcycleDataSheet.ModelYear,
					EngineDisplacement: mt.MotorcycleDataSheet.EngineDisplacement,
					TankCapacity:       mt.MotorcycleDataSheet.TankCapacity,
					Kilometers:         mt.MotorcycleDataSheet.Kilometers,
					Insurance: models.Insurance{
						IsActive:   mt.MotorcycleDataSheet.Insurance.IsActive,
						Company:    mt.MotorcycleDataSheet.Insurance.Company,
						PriceMoney: money,
						Currency:   mt.MotorcycleDataSheet.Insurance.Currency,
						Details:    mt.MotorcycleDataSheet.Insurance.Details,
						ExpireAt:   mt.MotorcycleDataSheet.Insurance.ExpireAt,
					},
				},
				FuelSupplies:   fromMongoFuelSuppliesToBlFuelSupplies(mt.FuelSupplies),
				Service:        fromMongoServiceToBlService(mt.Service),
				Inspection:     []models.Inspection{},
				AccidentReport: []models.AccidentReport{},
				CreatedAt:      mt.CreatedAt,
				UpdatedAt:      mt.UpdatedAt,
			})
		}

		return blMotorcycles
	}
	return nil
}

func fromBlFuelsSuppliesToMongoFuelsSupplies(blFuels []models.FuelSupplies) []models2.FuelSupplies {
	mongoFuels := make([]models2.FuelSupplies, 0)
	for _, f := range blFuels {
		mongoFuels = append(mongoFuels, f.ToMongoFuelSupplies())
	}
	return mongoFuels
}

func fromMongoFuelSuppliesToBlFuelSupplies(mongoFuels []models2.FuelSupplies) []models.FuelSupplies {
	blFuels := make([]models.FuelSupplies, 0)
	for _, mongoFuel := range mongoFuels {
		pricePerLitres, _ := money2.NewAmountFromFloat64(money2.EUR.Code(), mongoFuel.PetrolStation.FuelPricePerLitres)
		totalPrice, _ := money2.NewAmountFromFloat64(money2.EUR.Code(), mongoFuel.PetrolStation.TotalPrice)
		blFuels = append(blFuels, models.FuelSupplies{
			ID: mongoFuel.ID.Hex(),
			PetrolStation: models.PetrolStation{
				Name:               mongoFuel.PetrolStation.Name,
				Street:             mongoFuel.PetrolStation.Street,
				City:               mongoFuel.PetrolStation.City,
				Province:           mongoFuel.PetrolStation.Province,
				State:              mongoFuel.PetrolStation.State,
				FuelType:           models.FuelType(mongoFuel.PetrolStation.FuelType),
				FuelPricePerLitres: pricePerLitres,
				TotalLitres:        mongoFuel.PetrolStation.TotalLitres,
				TotalPrice:         totalPrice,
			},
			FullFuel:  mongoFuel.FullFuel,
			CreatedAt: mongoFuel.CreatedAt,
		})
	}
	return blFuels
}

func fromBlServicesToMongoServices(blServices []models.Service) []models2.Service {
	mongoServices := make([]models2.Service, 0)
	for _, f := range blServices {
		mongoServices = append(mongoServices, f.ToMongoService())
	}
	return mongoServices
}

func fromMongoServiceToBlService(mongoServices []models2.Service) []models.Service {
	blServices := make([]models.Service, 0)
	for _, s := range mongoServices {
		listOfDones := make([]models.ListOfDones, 0)
		for _, stuff := range s.ListOfDones {
			stuffPrice, _ := money2.NewAmountFromFloat64(money2.EUR.Code(), stuff.Price)
			listOfDones = append(listOfDones, models.ListOfDones{
				ID:    stuff.ID.Hex(),
				Name:  stuff.Name,
				Note:  stuff.Note,
				Price: stuffPrice,
			})
		}
		vatPrice, _ := money2.NewAmountFromFloat64(money2.EUR.Code(), s.VatPrice)
		totalPrice, _ := money2.NewAmountFromFloat64(money2.EUR.Code(), s.TotalPrice)
		manpowerPrice, _ := money2.NewAmountFromFloat64(money2.EUR.Code(), s.ManpowerPrice)
		blServices = append(blServices, models.Service{
			ID:   s.ID.Hex(),
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
		})
	}

	return blServices
}
