package server

import (
	"encoding/json"
	"errors"
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

	gotUser, gotUserErr := s.businessLogic.GetUserByUsername(getMotoInfo.Username)
	if gotUserErr != nil {
		err := map[string]interface{}{"MotorcyclesGetByLicensePlateRoute": gotUserErr.Error()}
		s.HandleRouteError(writer, err, http.StatusNotFound)
		return
	}

	if !gotUser.IsLoggedIn {
		err := map[string]interface{}{"MotorcyclesGetByLicensePlateRoute": errors.New("user is not logged in").Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	jwtValidationErr := s.ValidateJwtToken(token, gotUser.Token)
	if jwtValidationErr != nil {
		err := map[string]interface{}{"MotorcyclesGetByLicensePlateRoute": jwtValidationErr.Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
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
