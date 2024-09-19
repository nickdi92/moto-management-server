package server

import (
	"errors"
	"fmt"
	"moto-management-server/business_logic"
	"moto-management-server/utils"
	"net/http"

	"github.com/thedevsaddam/govalidator"
)

var RegisterRoute = func(s *MotoManagementServer, writer http.ResponseWriter, request *http.Request) {
	/**
	* 1. Get username and password
	* 2. Check if there isn't any other users with same username
	* 3. Create New user with new jwt token
	 */

	username := request.PostFormValue("username")
	password, _ := utils.Password(request.PostFormValue("password")).Hash()

	rules := govalidator.MapData{
		"username": []string{"required"},
		"password": []string{"required"},
	}
	validErr := s.ValidateRequest(request, rules)
	if validErr != nil {
		err := map[string]interface{}{"registerRouteErr": validErr}
		writer.WriteHeader(http.StatusUnauthorized)
		s.HandleRouteError(writer, err)
		return
	}

	findUser, _ := s.businessLogic.GetUserByUsername(username)
	if findUser.ID != "" {
		err := map[string]interface{}{"registerRouteErr": errors.New(fmt.Sprintf("user %s already exist", username)).Error()}
		writer.WriteHeader(http.StatusUnauthorized)
		s.HandleRouteError(writer, err)
		return
	}

	token := s.token.NewToken(username, password)
	tokenErr := token.GenerateToken()
	if tokenErr != nil {
		err := map[string]interface{}{"routeRegisterErr": tokenErr.Error()}
		writer.WriteHeader(http.StatusUnauthorized)
		s.HandleRouteError(writer, err)
		return
	}

	newUser := business_logic.User{
		Username:   username,
		Password:   password,
		Name:       request.PostFormValue("name"),
		Lastname:   request.PostFormValue("lastname"),
		Token:      token.Token,
		ExpireAt:   token.ExpiresAt,
		IsLoggedIn: false,
	}

	userCreated, userCreatedErr := s.businessLogic.CreateNewUser(newUser)
	if userCreatedErr != nil {
		err := map[string]interface{}{"registerRouteErr": userCreatedErr.Error()}
		writer.WriteHeader(http.StatusUnauthorized)
		s.HandleRouteError(writer, err)
		return
	}
	s.HandleResponse(writer, userCreated)
}
