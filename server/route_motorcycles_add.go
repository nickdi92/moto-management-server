package server

import (
	"encoding/json"
	"errors"
	"io"
	"moto-management-server/server/models"
	"net/http"
)

var MotorcyclesAddRoute = func(s *MotoManagementServer, writer http.ResponseWriter, request *http.Request) {
	var motorBiker models.AddMotorcycleRequest
	body, _ := io.ReadAll(request.Body)
	_ = json.Unmarshal(body, &motorBiker)

	validationErr := s.ValidateRequest(motorBiker)
	if validationErr != nil {
		err := map[string]interface{}{"MotorcyclesAddRoute": validationErr.Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	token, authValidationErr := s.ValidateAuthorization(writer, request)
	if token == "" && authValidationErr == nil {
		return
	}

	gotUser, gotUserErr := s.businessLogic.GetUserByUsername(motorBiker.Username)
	if gotUserErr != nil {
		err := map[string]interface{}{"MotorcyclesAddRoute": gotUserErr.Error()}
		s.HandleRouteError(writer, err, http.StatusNotFound)
		return
	}

	if !gotUser.IsLoggedIn {
		err := map[string]interface{}{"MotorcyclesAddRoute": errors.New("user is not logged in").Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	jwtValidationErr := s.ValidateJwtToken(token, gotUser.Token)
	if jwtValidationErr != nil {
		err := map[string]interface{}{"MotorcyclesAddRoute": jwtValidationErr.Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	blMotorBiker, blResponseErr := s.businessLogic.UpdateUser(fromServerMotorBikerToBlUSer(motorBiker))
	if blResponseErr != nil {
		err := map[string]interface{}{"MotorcyclesAddRoute": blResponseErr.Error()}
		s.HandleRouteError(writer, err, http.StatusInternalServerError)
		return
	}

	s.HandleResponse(writer, fromBlMotorBikerToServerMotorBiker(blMotorBiker))

}
