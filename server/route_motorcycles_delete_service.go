package server

import (
	"encoding/json"
	"errors"
	"io"
	"moto-management-server/server/models"
	"net/http"
)

var MotorcyclesDeleteServiceRoute = func(s *MotoManagementServer, writer http.ResponseWriter, request *http.Request) {
	var serviceToRemove models.DeleteServiceRequest
	body, _ := io.ReadAll(request.Body)
	_ = json.Unmarshal(body, &serviceToRemove)

	validationErr := s.ValidateRequest(serviceToRemove)
	if validationErr != nil {
		err := map[string]interface{}{"MotorcyclesDeleteServiceRoute": validationErr.Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	token, authValidationErr := s.ValidateAuthorization(writer, request)
	if token == "" && authValidationErr == nil {
		return
	}

	gotUser, gotUserErr := s.businessLogic.GetUserByUsername(serviceToRemove.Username)
	if gotUserErr != nil {
		err := map[string]interface{}{"MotorcyclesDeleteServiceRoute": gotUserErr.Error()}
		s.HandleRouteError(writer, err, http.StatusNotFound)
		return
	}

	if !gotUser.IsLoggedIn {
		err := map[string]interface{}{"MotorcyclesDeleteServiceRoute": errors.New("user is not logged in").Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	jwtValidationErr := s.ValidateJwtToken(token, gotUser.Token)
	if jwtValidationErr != nil {
		err := map[string]interface{}{"MotorcyclesDeleteServiceRoute": jwtValidationErr.Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	userUpdated, err := s.businessLogic.RemoveServiceFromMotorcycle(
		serviceToRemove.Username,
		serviceToRemove.LicensePlate,
		serviceToRemove.ServiceId,
	)

	if err != nil {
		respErr := map[string]interface{}{"MotorcyclesDeleteServiceRoute": err.Error()}
		s.HandleRouteError(writer, respErr, http.StatusInternalServerError)
		return
	}

	s.HandleResponse(writer, models.AddServiceToMotorcycleResponse{
		StatusCode: http.StatusOK,
		User:       fromBlUserToServerUser(userUpdated),
	})
}
