package server

import (
	"fmt"
	models2 "moto-management-server/business_logic/models"
	"moto-management-server/server/models"
	"time"

	money2 "github.com/Rhymond/go-money"
)

func fromUserRegisterRequestToBlUser(registerUser models.RegisterUserRequest) models2.User {
	return models2.User{
		Username:   registerUser.Username,
		Password:   registerUser.Password,
		Email:      registerUser.Email,
		Name:       registerUser.Name,
		Lastname:   registerUser.Lastname,
		IsLoggedIn: false,
	}
}

func fromBlUserToUserRegisterRequest(blUser models2.User) models.RegisterUserRequest {
	return models.RegisterUserRequest{
		Username:   blUser.Username,
		Password:   blUser.Password,
		Email:      blUser.Email,
		Name:       blUser.Name,
		Lastname:   blUser.Lastname,
		Token:      blUser.Token,
		ExpireAt:   blUser.ExpireAt,
		IsLoggedIn: false,
	}
}

/********/

func fromBlUserToUserLoginRequest(blUser models2.User) models.UserLoginRequest {
	return models.UserLoginRequest{
		Username:   blUser.Username,
		Password:   blUser.Password,
		Token:      blUser.Token,
		ExpireAt:   blUser.ExpireAt,
		IsLoggedIn: blUser.IsLoggedIn,
	}
}

func fromServerMotorBikerToBlUSer(biker models.MotorBiker) models2.User {
	blMotorBiker := models2.User{
		Username:    biker.Username,
		Motorcycles: fromServerMotorcyclesToBlMotorcycles(biker.Motorcycles),
	}

	return blMotorBiker
}

func fromBlMotorBikerToServerMotorBiker(biker models2.User) models.MotorBiker {
	serverBiker := models.MotorBiker{
		Username:    biker.Username,
		Motorcycles: fromBlMotorcyclesToServerMotorcycles(biker.Motorcycles),
	}

	return serverBiker
}

func fromBlMotorcyclesToServerMotorcycles(blMotorcycles []models2.Motorcycle) []models.Motorcycle {
	if blMotorcycles != nil && len(blMotorcycles) > 0 {
		serverMotorcycles := make([]models.Motorcycle, 0)
		for _, mt := range blMotorcycles {
			serverMotorcycles = append(serverMotorcycles, models.Motorcycle{
				ID:           mt.ID,
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
						PriceMoney: mt.MotorcycleDataSheet.Insurance.PriceMoney.AsMajorUnits(),
						Currency:   mt.MotorcycleDataSheet.Insurance.Currency,
						Details:    mt.MotorcycleDataSheet.Insurance.Details,
						ExpireAt:   mt.MotorcycleDataSheet.Insurance.ExpireAt.String(),
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

		return serverMotorcycles
	}
	return nil
}
func fromServerMotorcyclesToBlMotorcycles(serverMotorcycles []models.Motorcycle) []models2.Motorcycle {
	if serverMotorcycles != nil && len(serverMotorcycles) > 0 {
		blMotorcycles := make([]models2.Motorcycle, 0)
		for _, mt := range serverMotorcycles {
			currency := mt.MotorcycleDataSheet.Insurance.Currency
			if currency == "" {
				currency = money2.EUR
			}
			money := money2.NewFromFloat(mt.MotorcycleDataSheet.Insurance.PriceMoney, currency)
			expireAt, _ := time.Parse(time.DateOnly, mt.MotorcycleDataSheet.Insurance.ExpireAt)
			
			blMt := models2.Motorcycle{
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
						PriceMoney: money,
						Details:    mt.MotorcycleDataSheet.Insurance.Details,
						ExpireAt:   &expireAt,
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
				blMt.ID = mt.ID
			}

			blMotorcycles = append(blMotorcycles, blMt)
		}

		return blMotorcycles
	}
	return nil
}
func fromBlUserToServerUser(blUser models2.User) models.User {
	return models.User{
		ID:          blUser.ID,
		Username:    blUser.Username,
		Password:    blUser.Password,
		Email:       blUser.Email,
		Name:        blUser.Name,
		Lastname:    blUser.Lastname,
		Token:       blUser.Token,
		ExpireAt:    blUser.ExpireAt,
		IsLoggedIn:  blUser.IsLoggedIn,
		Motorcycles: fromBlMotorcyclesToServerMotorcycles(blUser.Motorcycles),
	}
}
