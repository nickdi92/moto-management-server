package server

import (
	"moto-management-server/utils"
	"net/http"
	"fmt"
)

var LoginRoute = func(s *MotoManagementServer, writer http.ResponseWriter, request *http.Request) {
	jwt := request.Header.Get("Authorization")
	if jwt == "" {
		err := map[string]interface{}{"loginRouteErr": fmt.Errorf("Missing authorization Bearer").Error()}
		writer.WriteHeader(http.StatusUnauthorized)
		s.HandleRouteError(writer, err)
		return
	}

	username := request.PostFormValue("username")
	password, _ := utils.Password(request.PostFormValue("password")).Hash()

	token := s.token.NewToken(username, password)
	token.Token = jwt[len("Bearer "):]
	validateErr := token.ValidateToken()

	if validateErr != nil {
		err := map[string]interface{}{"loginRouteErr": validateErr.Error()}
		writer.WriteHeader(http.StatusUnauthorized)
		s.HandleRouteError(writer, err)
		return
	}

	user, userErr := s.businessLogic.GetUserByUsername(username)
	if userErr != nil {
		err := map[string]interface{}{"loginRouteErr": userErr.Error()}
		writer.WriteHeader(http.StatusUnauthorized)
		s.HandleRouteError(writer, err)
		return
	}

	user.IsLoggedIn = true
	user, updateErr := s.businessLogic.UpdateUser(user)

	if updateErr != nil {
		err := map[string]interface{}{"loginRouteErr": updateErr.Error()}
		writer.WriteHeader(http.StatusUnauthorized)
		s.HandleRouteError(writer, err)
		return
	}

	s.HandleResponse(writer, user)
}
