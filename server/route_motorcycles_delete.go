package server

import (
	"encoding/json"
	"io"
	"moto-management-server/server/models"
	"net/http"
)

var MotorcyclesDeleteRoute = func(s *MotoManagementServer, writer http.ResponseWriter, request *http.Request) {
	var deleteMotorcycle models.DeleteMotorcycleRequest
	body, _ := io.ReadAll(request.Body)
	_ = json.Unmarshal(body, &deleteMotorcycle)

	validationErr := s.ValidateRequest(deleteMotorcycle)
	if validationErr != nil {
		err := map[string]interface{}{"motorcycleDeleteRoute": validationErr.Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	token, authValidationErr := s.ValidateAuthorization(writer, request)
	if token == "" && authValidationErr == nil {
		return
	}

	gotUser, gotUserErr := s.businessLogic.GetUserByUsername(deleteMotorcycle.Username)
	if gotUserErr != nil {
		err := map[string]interface{}{"MotorcyclesDeleteRoute": gotUserErr.Error()}
		s.HandleRouteError(writer, err, http.StatusNotFound)
		return
	}

	jwtValidationErr := s.ValidateJwtToken(token, gotUser.Token)
	if jwtValidationErr != nil {
		err := map[string]interface{}{"MotorcyclesDeleteRoute": jwtValidationErr.Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	isDeleted, deleteErr := s.businessLogic.DeleteMotorbike(deleteMotorcycle.Username, deleteMotorcycle.LicensePlate)
	if deleteErr != nil {
		err := map[string]interface{}{"motorcycleDeleteRoute": deleteErr.Error()}
		s.HandleRouteError(writer, err, http.StatusInternalServerError)
		return
	}

	s.HandleResponse(writer, models.DeleteMotorcycleResponse{StatusCode: http.StatusOK, IsDeleted: isDeleted})
}
