package server

import (
	"encoding/json"
	"errors"
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

	gotUser, gotUserErr := s.businessLogic.GetUserByUsername(userUpdateRequest.CreateUserRequest.Username)
	if gotUserErr != nil {
		err := map[string]interface{}{"UserUpdateRoute": gotUserErr.Error()}
		s.HandleRouteError(writer, err, http.StatusNotFound)
		return
	}

	if !gotUser.IsLoggedIn {
		err := map[string]interface{}{"UserUpdateRoute": errors.New("user is not logged in").Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	jwtValidationErr := s.ValidateJwtToken(token, gotUser.Token)
	if jwtValidationErr != nil {
		err := map[string]interface{}{"UserUpdateRoute": jwtValidationErr.Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	userUpdateRequest.CreateUserRequest.IsLoggedIn = gotUser.IsLoggedIn
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
