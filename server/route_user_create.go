package server

import (
	"encoding/json"
	"fmt"
	"io"
	"moto-management-server/server/models"
	"moto-management-server/utils"
	"moto-management-server/utils/validator"
	"net/http"
	"strings"
)

var UserCreateRoute = func(s *MotoManagementServer, writer http.ResponseWriter, request *http.Request) {
	/**
	* 1. Get username and password
	* 2. Check if there isn't any other users with same username
	* 3. Create New user with new jwt token
	 */

	var registerUserRequest models.CreateUserRequest
	body, _ := io.ReadAll(request.Body)
	_ = json.Unmarshal([]byte(body), &registerUserRequest)

	validationErr := s.ValidateRequest(registerUserRequest)
	if validationErr != nil {
		err := map[string]interface{}{"registerRouteErr": validationErr.Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	emailValidator := validator.NewEmailValidator(registerUserRequest.Email)
	validationEmailErr := emailValidator.Validate()
	if validationEmailErr != nil {
		err := map[string]interface{}{"registerRouteErr": validationEmailErr.Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	registerUserRequest.Password, _ = utils.Password(registerUserRequest.Password).Hash()
	registerUserRequest.Username = strings.ToLower(registerUserRequest.Username)

	findUser, _ := s.businessLogic.GetUserByUsername(registerUserRequest.Username)
	if findUser.ID != "" {
		err := map[string]interface{}{"registerRouteErr": fmt.Errorf("user %s already exist", registerUserRequest.Username).Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	token := s.token.NewToken(registerUserRequest.Username, registerUserRequest.Password)
	tokenErr := token.GenerateToken()
	if tokenErr != nil {
		err := map[string]interface{}{"routeRegisterErr": tokenErr.Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	newUser := fromUserRegisterRequestToBlUser(registerUserRequest)
	newUser.Token = token.Token
	newUser.ExpireAt = token.ExpiresAt

	_, userCreatedErr := s.businessLogic.CreateNewUser(newUser)
	if userCreatedErr != nil {
		err := map[string]interface{}{"registerRouteErr": userCreatedErr.Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}
	response := models.CreateUserResponse{
		StatusCode: http.StatusOK,
		Token:      token.Token,
		ExpireAt:   token.ExpiresAt,
	}
	s.HandleResponse(writer, response)
}
