package server

import (
	"encoding/json"
	"io"
	"moto-management-server/server/models"
	"net/http"
)

var MotorcyclesCreateRoute = func(s *MotoManagementServer, writer http.ResponseWriter, request *http.Request) {
	var motorcycle models.Motorcycle
	body, _ := io.ReadAll(request.Body)
	_ = json.Unmarshal(body, &motorcycle)

	validationErr := s.ValidateRequest(motorcycle)
	if validationErr != nil {
		err := map[string]interface{}{"motorcycleCreateRoute": validationErr.Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	token, authValidationErr := s.ValidateAuthorization(writer, request)
	if token == "" && authValidationErr == nil {
		return
	}

	user, userErr := s.businessLogic.GetUserByUsername(motorcycle.Username)
	if userErr != nil {
		err := map[string]interface{}{"motorcycleCreateRoute": userErr.Error()}
		s.HandleRouteError(writer, err, http.StatusInternalServerError)
		return
	}

	blResponse, blResponseErr := s.businessLogic.AddMotorcycleToUser(user, fromServerMotorcycleToBlMotorcycle(motorcycle))
	if blResponseErr != nil {
		err := map[string]interface{}{"motorcycleCreateRoute": blResponseErr.Error()}
		s.HandleRouteError(writer, err, http.StatusInternalServerError)
		return
	}

	s.HandleResponse(writer, fromBlMotorcycleToServerMotorcycle(blResponse))

}
