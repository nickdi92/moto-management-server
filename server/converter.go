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

func fromServerMotorcycleToBlMotorcycle(motorcycle models.Motorcycle) models2.Motorcycle {
	return models2.Motorcycle{}
}

func fromBlMotorcycleToServerMotorcycle(motorcycle models2.Motorcycle) models.Motorcycle {
	return models.Motorcycle{}
}
