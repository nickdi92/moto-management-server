package server

import (
	"encoding/json"
	"fmt"
	"io"
	"moto-management-server/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var LoginRoute = func(s *MotoManagementServer, writer http.ResponseWriter, request *http.Request) {
	jwt := request.Header.Get("Authorization")
	if jwt == "" {
		err := map[string]interface{}{"loginRouteErr": fmt.Errorf("missing authorization Bearer").Error()}
		writer.WriteHeader(http.StatusUnauthorized)
		s.HandleRouteError(writer, err)
		return
	}

	var userLogin UserLoginRequest
	bodyReader, _ := io.ReadAll(request.Body)
	_ = json.Unmarshal([]byte(bodyReader), &userLogin)

	validate := validator.New()
	validationErr := validate.Struct(userLogin)
	if validationErr != nil {
		err := map[string]interface{}{"loginRouteErr": validationErr.Error()}
		writer.WriteHeader(http.StatusUnauthorized)
		s.HandleRouteError(writer, err)
		return
	}

	userLogin.Password, _ = utils.Password(userLogin.Password).Hash()

	token := s.token.NewToken(userLogin.Username, userLogin.Password)
	token.Token = jwt[len("Bearer "):]
	validateErr := token.ValidateToken()

	if validateErr != nil {
		err := map[string]interface{}{"loginRouteErr": validateErr.Error()}
		writer.WriteHeader(http.StatusUnauthorized)
		s.HandleRouteError(writer, err)
		return
	}

	user, userErr := s.businessLogic.GetUserByUsername(userLogin.Username)
	if userErr != nil {
		err := map[string]interface{}{"loginRouteErr": userErr.Error()}
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
