package server

import (
	"encoding/json"
	"fmt"
	"io"
	"moto-management-server/utils"
	"moto-management-server/utils/validator"
	"net/http"
	"strings"
)

var RegisterRoute = func(s *MotoManagementServer, writer http.ResponseWriter, request *http.Request) {
	/**
	* 1. Get username and password
	* 2. Check if there isn't any other users with same username
	* 3. Create New user with new jwt token
	 */

	var registerUserRequest RegisterUserRequest
	body, _ := io.ReadAll(request.Body)
	_ = json.Unmarshal([]byte(body), &registerUserRequest)

	validationErr := s.ValidateRequest(registerUserRequest)
	if validationErr != nil {
		err := map[string]interface{}{"registerRouteErr": validationErr.Error()}
		writer.WriteHeader(http.StatusUnauthorized)
		s.HandleRouteError(writer, err)
		return
	}

	emailValidator := validator.NewEmailValidator(registerUserRequest.Email)
	validationEmailErr := emailValidator.Validate()
	if validationEmailErr != nil {
		err := map[string]interface{}{"registerRouteErr": validationEmailErr.Error()}
		writer.WriteHeader(http.StatusUnauthorized)
		s.HandleRouteError(writer, err)
		return
	}

	registerUserRequest.Password, _ = utils.Password(registerUserRequest.Password).Hash()
	registerUserRequest.Username = strings.ToLower(registerUserRequest.Username)

	findUser, _ := s.businessLogic.GetUserByUsername(registerUserRequest.Username)
	if findUser.ID != "" {
		err := map[string]interface{}{"registerRouteErr": fmt.Errorf("user %s already exist", registerUserRequest.Username).Error()}
		writer.WriteHeader(http.StatusUnauthorized)
		s.HandleRouteError(writer, err)
		return
	}

	token := s.token.NewToken(registerUserRequest.Username, registerUserRequest.Password)
	tokenErr := token.GenerateToken()
	if tokenErr != nil {
		err := map[string]interface{}{"routeRegisterErr": tokenErr.Error()}
		writer.WriteHeader(http.StatusUnauthorized)
		s.HandleRouteError(writer, err)
		return
	}

	newUser := fromUserRegisterRequestToBlUser(registerUserRequest)
	newUser.Token = token.Token
	newUser.ExpireAt = token.ExpiresAt

	userCreated, userCreatedErr := s.businessLogic.CreateNewUser(newUser)
	if userCreatedErr != nil {
		err := map[string]interface{}{"registerRouteErr": userCreatedErr.Error()}
		writer.WriteHeader(http.StatusUnauthorized)
		s.HandleRouteError(writer, err)
		return
	}
	s.HandleResponse(writer, fromBlUserToUserRegisterRequest(userCreated))
}
