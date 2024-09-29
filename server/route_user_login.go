package server

import (
	"encoding/json"
	"io"
	"moto-management-server/server/models"
	"moto-management-server/utils"
	"net/http"
)

var LoginRoute = func(s *MotoManagementServer, writer http.ResponseWriter, request *http.Request) {

	jwtToken, err := s.ValidateAuthorization(writer, request)
	if jwtToken == "" && err == nil {
		// It means there is an error
		return
	}

	var userLogin models.UserLoginRequest
	bodyReader, _ := io.ReadAll(request.Body)
	_ = json.Unmarshal([]byte(bodyReader), &userLogin)

	validationErr := s.ValidateRequest(userLogin)
	if validationErr != nil {
		err := map[string]interface{}{"loginRouteErr": validationErr.Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	userLogin.Password, _ = utils.Password(userLogin.Password).Hash()

	user, userErr := s.businessLogic.GetUserByUsername(userLogin.Username)
	if userErr != nil {
		err := map[string]interface{}{"loginRouteErr": userErr.Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	token := s.token.NewToken(userLogin.Username, userLogin.Password)
	token.Token = jwtToken
	validateErr := token.ValidateToken(user.Token)
	if validateErr != nil {
		err := map[string]interface{}{"loginRouteErr": validateErr.Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	user.IsLoggedIn = true
	token.RefreshToken()
	user.Token = token.Token
	user, updateErr := s.businessLogic.UpdateUser(user)

	if updateErr != nil {
		err := map[string]interface{}{"loginRouteErr": updateErr.Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return
	}

	s.HandleResponse(writer, models.UserLoginResponse{
		StatusCode: http.StatusOK,
		Token:      user.Token,
		ExpireAt:   user.ExpireAt,
		IsLoggedIn: user.IsLoggedIn,
	})
}
