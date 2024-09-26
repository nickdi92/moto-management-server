package business_logic

import (
	"moto-management-server/business_logic/models"
	models2 "moto-management-server/database/models"

	money2 "github.com/Rhymond/go-money"

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
	}

	return blUser
}

func fromBlUserToMongoUser(blUser models.User) models2.User {
	mongoUser := models2.User{
		Username:    blUser.Username,
		Name:        blUser.Name,
		Lastname:    blUser.Lastname,
		Password:    blUser.Password,
		Email:       blUser.Email,
		Token:       blUser.Token,
		IsLoggedIn:  blUser.IsLoggedIn,
		Motorcycles: fromBlMotorcyclesToMongoMotorcycles(blUser.Motorcycles),
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
						PriceMoney: mt.MotorcycleDataSheet.Insurance.PriceMoney.AsMajorUnits(),
						Currency:   mt.MotorcycleDataSheet.Insurance.PriceMoney.Currency().Code,
						Details:    mt.MotorcycleDataSheet.Insurance.Details,
						ExpireAt:   mt.MotorcycleDataSheet.Insurance.ExpireAt,
					},
				},
				FuelSupplies:   models2.FuelSupplies{},
				Service:        models2.Service{},
				Inspection:     models2.Inspection{},
				AccidentReport: models2.AccidentReport{},
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
			money := money2.NewFromFloat(mt.MotorcycleDataSheet.Insurance.PriceMoney, mt.MotorcycleDataSheet.Insurance.Currency)
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
				FuelSupplies:   models.FuelSupplies{},
				Service:        models.Service{},
				Inspection:     models.Inspection{},
				AccidentReport: models.AccidentReport{},
				CreatedAt:      mt.CreatedAt,
				UpdatedAt:      mt.UpdatedAt,
			})
		}

		return blMotorcycles
	}
	return nil
}
