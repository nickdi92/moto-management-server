package server

import (
	"encoding/json"
	"io"
	"moto-management-server/utils"
	"net/http"
)

var LoginRoute = func(s *MotoManagementServer, writer http.ResponseWriter, request *http.Request) {

	jwtToken, err := s.ValidateAuthorization(writer, request)
	if jwtToken == "" && err == nil {
		// It means there is an error
		return
	}

	var userLogin UserLoginRequest
	bodyReader, _ := io.ReadAll(request.Body)
	_ = json.Unmarshal([]byte(bodyReader), &userLogin)

	validationErr := s.ValidateRequest(userLogin)
	if validationErr != nil {
		err := map[string]interface{}{"loginRouteErr": validationErr.Error()}
		writer.WriteHeader(http.StatusUnauthorized)
		s.HandleRouteError(writer, err)
		return
	}

	userLogin.Password, _ = utils.Password(userLogin.Password).Hash()

	user, userErr := s.businessLogic.GetUserByUsername(userLogin.Username)
	if userErr != nil {
		err := map[string]interface{}{"loginRouteErr": userErr.Error()}
		writer.WriteHeader(http.StatusUnauthorized)
		s.HandleRouteError(writer, err)
		return
	}

	token := s.token.NewToken(userLogin.Username, userLogin.Password)
	token.Token = jwtToken[len("Bearer "):]
	validateErr := token.ValidateToken(user.Token)
	if validateErr != nil {
		err := map[string]interface{}{"loginRouteErr": validateErr.Error()}
		writer.WriteHeader(http.StatusUnauthorized)
		s.HandleRouteError(writer, err)
		return
	}

	user.IsLoggedIn = true
	token.RefreshToken()
	user.Token = token.Token
	user, updateErr := s.businessLogic.UpdateUser(user)

	if updateErr != nil {
		err := map[string]interface{}{"loginRouteErr": updateErr.Error()}
		writer.WriteHeader(http.StatusUnauthorized)
		s.HandleRouteError(writer, err)
		return
	}

	s.HandleResponse(writer, fromBlUserToUserLoginRequest(user))
}
