package server

import (
	models2 "moto-management-server/business_logic/models"
	"moto-management-server/server/models"
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
		Username: biker.Username,
	}

	if biker.Motorcycles != nil && len(biker.Motorcycles) > 0 {
		blMotorcycles := make([]models2.Motorcycle, 0)
		for _, mt := range biker.Motorcycles {
			blMotorcycles = append(blMotorcycles, models2.Motorcycle{
				ID:           mt.ID,
				LicensePlate: mt.LicensePlate,
				MotorcycleDataSheet: models2.MotorcycleDataSheet{
					Name:               mt.MotorcycleDataSheet.Name,
					Model:              mt.MotorcycleDataSheet.Model,
					ModelYear:          mt.MotorcycleDataSheet.ModelYear,
					EngineDisplacement: mt.MotorcycleDataSheet.EngineDisplacement,
					TankCapacity:       mt.MotorcycleDataSheet.TankCapacity,
					Insurance: models2.Insurance{
						IsActive:   mt.MotorcycleDataSheet.Insurance.IsActive,
						Company:    mt.MotorcycleDataSheet.Insurance.Company,
						PriceMoney: mt.MotorcycleDataSheet.Insurance.PriceMoney,
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
			})
		}

		blMotorBiker.Motorcycles = blMotorcycles
	}

	return blMotorBiker
}

func fromBlMotorBikerToServerMotorBiker(biker models2.User) models.MotorBiker {
	serverBiker := models.MotorBiker{
		Username: biker.Username,
	}

	if biker.Motorcycles != nil && len(biker.Motorcycles) > 0 {
		serverMotorcycles := make([]models.Motorcycle, 0)
		for _, mt := range biker.Motorcycles {
			serverMotorcycles = append(serverMotorcycles, models.Motorcycle{
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

		serverBiker.Motorcycles = serverMotorcycles
	}

	return serverBiker
}
