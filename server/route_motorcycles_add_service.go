package server

import (
	"encoding/json"
	"io"
	"moto-management-server/server/models"
	"net/http"
)

var MotorcyclesAddServiceRoute = func(s *MotoManagementServer, writer http.ResponseWriter, request *http.Request) {
	var serviceToAdd models.AddServiceToMotorcycleRequest
	body, _ := io.ReadAll(request.Body)
	_ = json.Unmarshal(body, &serviceToAdd)

	validationErr := s.ValidateRequest(serviceToAdd)
	if validationErr != nil {
		err := map[string]interface{}{"MotorcyclesAddServiceRoute": validationErr.Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	token, authValidationErr := s.ValidateAuthorization(writer, request)
	if token == "" && authValidationErr == nil {
		return
	}

	gotUser, gotUserErr := s.businessLogic.GetUserByUsername(serviceToAdd.Username)
	if gotUserErr != nil {
		err := map[string]interface{}{"MotorcyclesAddServiceRoute": gotUserErr.Error()}
		s.HandleRouteError(writer, err, http.StatusNotFound)
		return
	}

	jwtValidationErr := s.ValidateJwtToken(token, gotUser.Token)
	if jwtValidationErr != nil {
		err := map[string]interface{}{"MotorcyclesAddServiceRoute": jwtValidationErr.Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	userUpdated, err := s.businessLogic.AddServiceToMotorcycle(
		serviceToAdd.Username,
		serviceToAdd.LicensePlate,
		serviceToAdd.Service.ToBusinessLogicModel(),
	)

	if err != nil {
		respErr := map[string]interface{}{"MotorcyclesAddServiceRoute": err.Error()}
		s.HandleRouteError(writer, respErr, http.StatusInternalServerError)
		return
	}

	s.HandleResponse(writer, models.AddServiceToMotorcycleResponse{
		StatusCode: http.StatusOK,
		User:       fromBlUserToServerUser(userUpdated),
	})
}
