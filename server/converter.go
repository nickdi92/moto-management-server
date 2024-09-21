package server

import "moto-management-server/business_logic"

func fromUserRegisterRequestToBlUser(registerUser RegisterUserRequest) business_logic.User {
	return business_logic.User{
		Username:   registerUser.Username,
		Password:   registerUser.Password,
		Email:      registerUser.Email,
		Name:       registerUser.Name,
		Lastname:   registerUser.Lastname,
		IsLoggedIn: false,
	}
}

func fromBlUserToUserRegisterRequest(blUser business_logic.User) RegisterUserRequest {
	return RegisterUserRequest{
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

func fromBlUserToUserLoginRequest(blUser business_logic.User) UserLoginRequest {
	return UserLoginRequest{
		Username:   blUser.Username,
		Password:   blUser.Password,
		Token:      blUser.Token,
		ExpireAt:   blUser.ExpireAt,
		IsLoggedIn: blUser.IsLoggedIn,
	}
}
