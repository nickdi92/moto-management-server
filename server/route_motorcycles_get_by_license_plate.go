package server

import (
	"encoding/json"
	"io"
	"moto-management-server/server/models"
	"net/http"
)

var MotorcyclesGetByLicensePlateRoute = func(s *MotoManagementServer, writer http.ResponseWriter, request *http.Request) {
	var getMotoInfo models.GetMotorcycleByLicensePlateRequest
	body, _ := io.ReadAll(request.Body)
	_ = json.Unmarshal(body, &getMotoInfo)

	validationErr := s.ValidateRequest(getMotoInfo)
	if validationErr != nil {
		err := map[string]interface{}{"MotorcyclesGetByLicensePlateRoute": validationErr.Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	token, authValidationErr := s.ValidateAuthorization(writer, request)
	if token == "" && authValidationErr == nil {
		return
	}

	motorcycle, getMotoErr := s.businessLogic.GetMotorcycleByLicensePlate(getMotoInfo.Username, getMotoInfo.LicensePlate)
	if getMotoErr != nil {
		err := map[string]interface{}{"MotorcyclesGetByLicensePlateRoute": getMotoErr.Error()}
		s.HandleRouteError(writer, err, http.StatusInternalServerError)
		return
	}

	response := models.GetMotorcycleByLicensePlateResponse{
		StatusCode: http.StatusOK,
		Motorcycle: fromBlMotoToServerMoto(motorcycle),
	}
	s.HandleResponse(writer, response)
}
