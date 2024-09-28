package server

import (
	"encoding/json"
	"io"
	"moto-management-server/server/models"
	"net/http"
)

var MotorcyclesAddFuelRoute = func(s *MotoManagementServer, writer http.ResponseWriter, request *http.Request) {
	var fuel models.AddFuelToMotorcycleRequest
	body, _ := io.ReadAll(request.Body)
	_ = json.Unmarshal(body, &fuel)

	validationErr := s.ValidateRequest(fuel)
	if validationErr != nil {
		err := map[string]interface{}{"MotorcyclesAddFuelRoute": validationErr.Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	token, authValidationErr := s.ValidateAuthorization(writer, request)
	if token == "" && authValidationErr == nil {
		return
	}

	motorcycle, err := s.businessLogic.AddFuelToMotorcycle(
		fuel.Username,
		fuel.LicensePlate,
		fuel.FuelSupplies.ToBusinessLogicModel(),
	)

	if err != nil {
		respErr := map[string]interface{}{"MotorcyclesAddFuelRoute": err.Error()}
		s.HandleRouteError(writer, respErr, http.StatusInternalServerError)
		return
	}

	s.HandleResponse(writer, models.AddFuelToMotorcycleResponse{
		StatusCode: http.StatusOK,
		Motorcycle: fromBlMotoToServerMoto(motorcycle),
	})
}
