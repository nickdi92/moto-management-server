package server

import (
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"moto-management-server/utils"
	"net/http"
)

var AuthRoute = func(s *MotoManagementServer, writer http.ResponseWriter, request *http.Request) {
	username := request.FormValue("username")
	password := request.FormValue("password")

	rules := govalidator.MapData{
		"username": []string{"required"},
		"password": []string{"required"},
	}

	opts := govalidator.Options{
		Request: request,
		Rules:   rules,
	}

	v := govalidator.New(opts)
	validationErr := v.Validate()
	if len(validationErr) > 0 {
		err := map[string]interface{}{"validationErr": validationErr}
		s.HandleRouteError(writer, err)
	} else {
		findUser, findUserErr := s.businessLogic.GetUserByUsername(username)
		if findUserErr != nil {
			err := map[string]interface{}{"findUserErr": findUserErr}
			s.HandleRouteError(writer, err)
		} else {
			s.HandleResponse(writer, findUser)
		}
	}

	utils.SuccessOutput(fmt.Sprintf("username %s with psw %s", username, password))
}
