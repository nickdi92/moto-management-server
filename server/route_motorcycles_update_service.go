package server

import (
	"encoding/json"
	"errors"
	"io"
	"moto-management-server/server/models"
	"net/http"
)

var MotorcyclesUpdateServiceRoute = func(s *MotoManagementServer, writer http.ResponseWriter, request *http.Request) {
	var serviceToUpdate models.UpdateServiceRequest
	body, _ := io.ReadAll(request.Body)
	_ = json.Unmarshal(body, &serviceToUpdate)

	validationErr := s.ValidateRequest(serviceToUpdate)
	if validationErr != nil {
		err := map[string]interface{}{"MotorcyclesUpdateServiceRoute": validationErr.Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	token, authValidationErr := s.ValidateAuthorization(writer, request)
	if token == "" && authValidationErr == nil {
		return
	}

	gotUser, gotUserErr := s.businessLogic.GetUserByUsername(serviceToUpdate.Username)
	if gotUserErr != nil {
		err := map[string]interface{}{"MotorcyclesUpdateServiceRoute": gotUserErr.Error()}
		s.HandleRouteError(writer, err, http.StatusNotFound)
		return
	}

	if !gotUser.IsLoggedIn {
		err := map[string]interface{}{"MotorcyclesUpdateServiceRoute": errors.New("user is not logged in").Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	jwtValidationErr := s.ValidateJwtToken(token, gotUser.Token)
	if jwtValidationErr != nil {
		err := map[string]interface{}{"MotorcyclesUpdateServiceRoute": jwtValidationErr.Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	updatedUser, updateErr := s.businessLogic.UpdateServiceToMotorcycle(
		serviceToUpdate.Username,
		serviceToUpdate.LicensePlate,
		serviceToUpdate.Service.ToBusinessLogicModel(),
	)

	if updateErr != nil {
		respErr := map[string]interface{}{"MotorcyclesUpdateServiceRoute": updateErr.Error()}
		s.HandleRouteError(writer, respErr, http.StatusInternalServerError)
		return
	}

	s.HandleResponse(writer, models.UpdateServiceToMotorcycleResponse{
		StatusCode: http.StatusOK,
		User:       fromBlUserToServerUser(updatedUser),
	})
}
