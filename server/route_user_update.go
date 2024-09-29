package server

import (
	"encoding/json"
	"io"
	"moto-management-server/server/models"
	"net/http"
)

var UserUpdateRoute = func(s *MotoManagementServer, writer http.ResponseWriter, request *http.Request) {
	var userUpdateRequest models.UpdateUserRequest
	body, _ := io.ReadAll(request.Body)
	_ = json.Unmarshal(body, &userUpdateRequest)

	validationErr := s.ValidateRequest(userUpdateRequest)
	if validationErr != nil {
		err := map[string]interface{}{"UserUpdateRoute": validationErr.Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	token, authValidationErr := s.ValidateAuthorization(writer, request)
	if token == "" && authValidationErr == nil {
		return
	}

	updatedUser, updateErr := s.businessLogic.UpdateUser(fromUserRegisterRequestToBlUser(userUpdateRequest.CreateUserRequest))
	if updateErr != nil {
		err := map[string]interface{}{"UserUpdateRoute": updateErr.Error()}
		s.HandleRouteError(writer, err, http.StatusInternalServerError)
		return
	}

	s.HandleResponse(writer, models.UpdateUserResponse{
		StatusCode: http.StatusOK,
		User:       fromBlUserToServerUser(updatedUser),
	})

}
