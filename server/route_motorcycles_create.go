package server

import (
	"encoding/json"
	"io"
	"moto-management-server/server/models"
	"net/http"
)

var MotorcyclesCreateRoute = func(s *MotoManagementServer, writer http.ResponseWriter, request *http.Request) {
	var motorBiker models.MotorBiker
	body, _ := io.ReadAll(request.Body)
	_ = json.Unmarshal(body, &motorBiker)

	validationErr := s.ValidateRequest(motorBiker)
	if validationErr != nil {
		err := map[string]interface{}{"motorcycleCreateRoute": validationErr.Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	token, authValidationErr := s.ValidateAuthorization(writer, request)
	if token == "" && authValidationErr == nil {
		return
	}

	_, userErr := s.businessLogic.GetUserByUsername(motorBiker.Username)
	if userErr != nil {
		err := map[string]interface{}{"motorcycleCreateRoute": userErr.Error()}
		s.HandleRouteError(writer, err, http.StatusInternalServerError)
		return
	}

	blMotorBiker, blResponseErr := s.businessLogic.UpdateUser(fromServerMotorBikerToBlUSer(motorBiker))
	if blResponseErr != nil {
		err := map[string]interface{}{"motorcycleCreateRoute": blResponseErr.Error()}
		s.HandleRouteError(writer, err, http.StatusInternalServerError)
		return
	}

	s.HandleResponse(writer, fromBlMotorBikerToServerMotorBiker(blMotorBiker))

}
