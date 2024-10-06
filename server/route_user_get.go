package server

import (
	"encoding/json"
	"fmt"
	"io"
	"moto-management-server/server/models"
	"net/http"
)

var GetUserRoute = func(s *MotoManagementServer, writer http.ResponseWriter, request *http.Request) {
	var getUser models.GetUserRequest
	body, _ := io.ReadAll(request.Body)
	_ = json.Unmarshal(body, &getUser)
	validationErr := s.ValidateRequest(getUser)
	if validationErr != nil {
		err := map[string]interface{}{"getUserRoute": validationErr.Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	token, authValidationErr := s.ValidateAuthorization(writer, request)
	if token == "" && authValidationErr == nil {
		return
	}

	gotUser, gotUserErr := s.businessLogic.GetUserByUsername(getUser.Username)
	if gotUserErr != nil {
		err := map[string]interface{}{"GetUserRoute": gotUserErr.Error()}
		s.HandleRouteError(writer, err, http.StatusNotFound)
		return
	}

	if !gotUser.IsLoggedIn {
		err := map[string]interface{}{"GetUserRoute": fmt.Errorf("user is not logged in")}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	jwtValidationErr := s.ValidateJwtToken(token, gotUser.Token)
	if jwtValidationErr != nil {
		err := map[string]interface{}{"GetUserRoute": jwtValidationErr.Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	resp := models.GetUserResponse{
		StatusCode: http.StatusOK,
		User:       fromBlUserToServerUser(gotUser),
	}
	s.HandleResponse(writer, resp)
}
