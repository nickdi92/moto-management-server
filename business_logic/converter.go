package business_logic

import (
	"moto-management-server/business_logic/models"
	"moto-management-server/database"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func fromMongoUserToBlUser(mongoUser database.User) models.User {
	id := mongoUser.ID.Hex()
	if mongoUser.ID.IsZero() {
		id = ""
	}
	blUser := models.User{
		ID:         id,
		Username:   mongoUser.Username,
		Name:       mongoUser.Name,
		Lastname:   mongoUser.Lastname,
		Password:   mongoUser.Password,
		Email:      mongoUser.Email,
		Token:      mongoUser.Token,
		ExpireAt:   &mongoUser.ExpireAt,
		CreatedAt:  &mongoUser.CreatedAt,
		UpdatedAt:  &mongoUser.UpdatedAt,
		IsLoggedIn: mongoUser.IsLoggedIn,
	}

	if mongoUser.Motorcycles != nil {
		blMotorcycles := make([]models.Motorcycle, 0)
		for _, mt := range mongoUser.Motorcycles {
			blMotorcycles = append(blMotorcycles, models.Motorcycle{
				ID:           mt.ID,
				LicensePlate: mt.LicensePlate,
				MotorcycleDataSheet: models.MotorcycleDataSheet{
					Name:               mt.MotorcycleDataSheet.Name,
					Model:              mt.MotorcycleDataSheet.Model,
					ModelYear:          mt.MotorcycleDataSheet.ModelYear,
					EngineDisplacement: mt.MotorcycleDataSheet.EngineDisplacement,
					TankCapacity:       mt.MotorcycleDataSheet.TankCapacity,
					Insurance: models.Insurance{
						IsActive:   mt.MotorcycleDataSheet.Insurance.IsActive,
						Company:    mt.MotorcycleDataSheet.Insurance.Company,
						PriceMoney: mt.MotorcycleDataSheet.Insurance.PriceMoney,
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

		blUser.Motorcycles = blMotorcycles
	}

	return blUser
}

func fromBlUserToMongoUser(blUser models.User) database.User {
	mongoUser := database.User{
		Username:   blUser.Username,
		Name:       blUser.Name,
		Lastname:   blUser.Lastname,
		Password:   blUser.Password,
		Email:      blUser.Email,
		Token:      blUser.Token,
		IsLoggedIn: blUser.IsLoggedIn,
	}

	if blUser.ExpireAt != nil {
		mongoUser.ExpireAt = *blUser.ExpireAt
	}

	if blUser.ID != "" {
		mongoUser.ID, _ = primitive.ObjectIDFromHex(blUser.ID)
	}

	if blUser.Motorcycles != nil {
		mongoMotorcycles := make([]database.Motorcycle, 0)
		for _, mt := range blUser.Motorcycles {
			mongoMotorcycles = append(mongoMotorcycles, database.Motorcycle{
				ID:           mt.ID,
				LicensePlate: mt.LicensePlate,
				MotorcycleDataSheet: database.MotorcycleDataSheet{
					Name:               mt.MotorcycleDataSheet.Name,
					Model:              mt.MotorcycleDataSheet.Model,
					ModelYear:          mt.MotorcycleDataSheet.ModelYear,
					EngineDisplacement: mt.MotorcycleDataSheet.EngineDisplacement,
					TankCapacity:       mt.MotorcycleDataSheet.TankCapacity,
					Insurance: database.Insurance{
						IsActive:   mt.MotorcycleDataSheet.Insurance.IsActive,
						Company:    mt.MotorcycleDataSheet.Insurance.Company,
						PriceMoney: mt.MotorcycleDataSheet.Insurance.PriceMoney,
						Details:    mt.MotorcycleDataSheet.Insurance.Details,
						ExpireAt:   mt.MotorcycleDataSheet.Insurance.ExpireAt,
					},
				},
				FuelSupplies:   database.FuelSupplies{},
				Service:        database.Service{},
				Inspection:     database.Inspection{},
				AccidentReport: database.AccidentReport{},
				CreatedAt:      mt.CreatedAt,
				UpdatedAt:      mt.UpdatedAt,
			})
		}

		mongoUser.Motorcycles = mongoMotorcycles
	}
	return mongoUser
}
