package server

import (
	"encoding/json"
	"github.com/thedevsaddam/govalidator"
	"moto-management-server/utils"
	"net/http"
)

var AuthRoute = func(s *MotoManagementServer, writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	/**
	 * -> Request has username and password. Are mandatory
	 * -> If there is a user in the DB, check his Token
	 * -> If token is expired, generate new one, associate it to user
	 * -> return the token (the saved one or the new one)
	 */

	username := request.PostFormValue("username")
	password, _ := utils.Password(request.PostFormValue("password")).Hash()

	rules := govalidator.MapData{
		"username": []string{"required"},
		"password": []string{"required"},
	}
	validErr := s.ValidateRequest(request, rules)
	if validErr != nil {
		err := map[string]interface{}{"validationError": validErr}
		writer.WriteHeader(http.StatusUnauthorized)
		s.HandleRouteError(writer, err)
		return
	}

	user, checkUserErr := s.businessLogic.GetUserByUsername(username)
	if checkUserErr != nil {
		err := map[string]interface{}{"checkUserErr": checkUserErr.Error()}
		writer.WriteHeader(http.StatusNotFound)
		s.HandleRouteError(writer, err)
		return
	}

	if user.Password != password {
		err := map[string]interface{}{"validationError": "passwords do not match"}
		writer.WriteHeader(http.StatusUnauthorized)
		s.HandleRouteError(writer, err)
		return
	}

	token := s.token.NewToken(username, password)
	_ = token.GenerateToken()
	isValid := token.ValidateToken()
	if isValid != nil {
		err := map[string]interface{}{"JWTValidationError": isValid.Error()}
		writer.WriteHeader(http.StatusInternalServerError)
		s.HandleRouteError(writer, err)
	}

	user.Token = token.Token
	user.ExpireAt = token.ExpiresAt

	s.HandleResponse(writer, json.NewEncoder(writer).Encode(map[string]interface{}{
		"token": token,
	}))
}
