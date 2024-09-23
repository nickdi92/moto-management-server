package server

import (
	"encoding/json"
	"io"
	"moto-management-server/server/models"
	"net/http"
)

var MotorcyclesDeleteRoute = func(s *MotoManagementServer, writer http.ResponseWriter, request *http.Request) {
	var deleteMotorcycle models.DeleteMotorcycle
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

	isDeleted, deleteErr := s.businessLogic.DeleteMotorbike(deleteMotorcycle.Username, deleteMotorcycle.LicensePlate)
	if deleteErr != nil {
		err := map[string]interface{}{"motorcycleDeleteRoute": deleteErr.Error()}
		s.HandleRouteError(writer, err, http.StatusInternalServerError)
		return
	}

	s.HandleResponse(writer, models.DeleteMotorcycleResponse{IsDeleted: isDeleted})
}
