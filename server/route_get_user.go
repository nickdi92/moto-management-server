package server

import (
	"encoding/json"
	"io"
	"moto-management-server/server/models"
	"net/http"
)

var GetUserRoute = func(s *MotoManagementServer, writer http.ResponseWriter, request *http.Request) {
	var getUser models.GetUserRoute
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

	user, getUserErr := s.businessLogic.GetUserByUsername(getUser.Username)
	if getUserErr != nil {
		err := map[string]interface{}{"getUserRoute": getUserErr.Error()}
		s.HandleRouteError(writer, err, http.StatusInternalServerError)
		return
	}

	s.HandleResponse(writer, fromBlUserToServerUser(user))
}
