package server

import (
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

var AuthRoute = func(s *MotoManagementServer, writer http.ResponseWriter, request *http.Request) {
	/**
	 * -> Request has username and password. Are mandatory
	 * -> If there is a user in the DB, check his Token
	 * -> If token is expired, generate new one, associate it to user
	 * -> return the token (the saved one or the new one)
	 */

	username := request.PostFormValue("username")
	password := request.PostFormValue("password")
	rules := govalidator.MapData{
		"username": []string{"required"},
		"password": []string{"required"},
	}
	validErr := s.ValidateRequest(request, rules)
	if validErr != nil {
		err := map[string]interface{}{"validationError": validErr}
		s.HandleRouteError(writer, err)
		return
	}

	checkUser, checkUserErr := s.businessLogic.GetUserByUsername(username)
	if checkUserErr != nil {
		err := map[string]interface{}{"checkUserErr": checkUserErr.Error()}
		s.HandleRouteError(writer, err)
		return
	}

	jwtErr := s.ValidateJwtToken()
}
