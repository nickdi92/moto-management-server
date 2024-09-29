package server

import (
	models2 "moto-management-server/business_logic/models"
	"moto-management-server/server/models"
	"net/http"
	"time"

	money2 "github.com/govalues/money"
)

func fromUserRegisterRequestToBlUser(registerUser models.CreateUserRequest) models2.User {
	return models2.User{
		Username:   registerUser.Username,
		Password:   registerUser.Password,
		Email:      registerUser.Email,
		Name:       registerUser.Name,
		Lastname:   registerUser.Lastname,
		IsLoggedIn: false,
		Address: models2.Address{
			City:     registerUser.Address.City,
			Street:   registerUser.Address.Street,
			ZipCode:  registerUser.Address.ZipCode,
			Province: registerUser.Address.Province,
			State:    registerUser.Address.State,
		},
		UserRegistry: models2.Registry{
			FiscalCode: registerUser.UserRegistry.FiscalCode,
			DOB:        registerUser.UserRegistry.DOB,
		},
	}
}

func fromServerMotorBikerToBlUSer(biker models.AddMotorcycleRequest) models2.User {
	blMotorBiker := models2.User{
		Username:    biker.Username,
		Motorcycles: fromServerMotorcyclesToBlMotorcycles(biker.Motorcycles),
	}

	return blMotorBiker
}

func fromBlMotorBikerToServerMotorBiker(biker models2.User) models.AddMotorcycleResponse {
	serverBiker := models.AddMotorcycleResponse{
		StatusCode:  http.StatusOK,
		Motorcycles: fromBlMotorcyclesToServerMotorcycles(biker.Motorcycles),
	}

	return serverBiker
}

func fromBlMotorcyclesToServerMotorcycles(blMotorcycles []models2.Motorcycle) []models.Motorcycle {
	if blMotorcycles != nil {
		serverMotorcycles := make([]models.Motorcycle, 0)
		for _, mt := range blMotorcycles {
			serverMotorcycles = append(serverMotorcycles, fromBlMotoToServerMoto(mt))
		}

		return serverMotorcycles
	}
	return nil
}

func fromBlMotoToServerMoto(mt models2.Motorcycle) models.Motorcycle {
	if mt.LicensePlate == "" {
		return models.Motorcycle{}
	}
	priceMoney, _ := mt.MotorcycleDataSheet.Insurance.PriceMoney.Float64()
	return models.Motorcycle{
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
				PriceMoney: priceMoney,
				Currency:   mt.MotorcycleDataSheet.Insurance.Currency,
				Details:    mt.MotorcycleDataSheet.Insurance.Details,
				ExpireAt:   mt.MotorcycleDataSheet.Insurance.ExpireAt.String(),
			},
		},
		FuelSupplies:   fromBlFuelSuppliesToServerFuelSupplies(mt.FuelSupplies),
		Service:        fromBlServicesToServerServices(mt.Service),
		Inspection:     []models.Inspection{},
		AccidentReport: []models.AccidentReport{},
		CreatedAt:      mt.CreatedAt,
		UpdatedAt:      mt.UpdatedAt,
	}
}

func fromServerMotorcyclesToBlMotorcycles(serverMotorcycles []models.Motorcycle) []models2.Motorcycle {
	if serverMotorcycles != nil {
		blMotorcycles := make([]models2.Motorcycle, 0)
		for _, mt := range serverMotorcycles {
			currency := mt.MotorcycleDataSheet.Insurance.Currency
			if currency == "" {
				currency = money2.EUR.Code()
			}

			money, _ := money2.NewAmountFromFloat64(currency, mt.MotorcycleDataSheet.Insurance.PriceMoney)
			expireAt, _ := time.Parse(time.DateOnly, mt.MotorcycleDataSheet.Insurance.ExpireAt)

			fuels := make([]models2.FuelSupplies, 0)
			for _, f := range mt.FuelSupplies {
				fuels = append(fuels, f.ToBusinessLogicModel())
			}

			services := make([]models2.Service, 0)
			for _, s := range mt.Service {
				services = append(services, s.ToBusinessLogicModel())
			}

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
				FuelSupplies:   fuels,
				Service:        services,
				Inspection:     []models2.Inspection{},
				AccidentReport: []models2.AccidentReport{},
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
		Address: models.Address{
			City:     blUser.Address.City,
			Street:   blUser.Address.Street,
			ZipCode:  blUser.Address.ZipCode,
			Province: blUser.Address.Province,
			State:    blUser.Address.State,
		},
		UserRegistry: models.Registry{
			FiscalCode: blUser.UserRegistry.FiscalCode,
			DOB:        blUser.UserRegistry.DOB,
		},
	}
}

func fromBlFuelSuppliesToServerFuelSupplies(blFuel []models2.FuelSupplies) []models.FuelSupplies {
	serverFuels := make([]models.FuelSupplies, 0)
	for _, bf := range blFuel {
		var sf models.FuelSupplies
		serverFuels = append(serverFuels, sf.ToServerModel(bf))
	}
	return serverFuels
}

func fromBlServicesToServerServices(blServices []models2.Service) []models.Service {
	serverServices := make([]models.Service, 0)
	for _, bs := range blServices {
		var serverService models.Service
		serverServices = append(serverServices, serverService.ToServerModel(bs))
	}
	return serverServices
}
