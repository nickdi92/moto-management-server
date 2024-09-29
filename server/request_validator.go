package server

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func (s *MotoManagementServer) ValidateRequest(structToValidate interface{}) error {
	validate := validator.New()
	validationErr := validate.Struct(structToValidate)
	if validationErr != nil {
		return validationErr
	}
	return nil
}

func (s *MotoManagementServer) ValidateJwtToken(incomingJwtToken string, oldToken string) error {
	token := s.token.NewToken("", "")
	token.Token = incomingJwtToken
	return token.ValidateToken(oldToken)
}

func (s *MotoManagementServer) ValidateAuthorization(writer http.ResponseWriter, request *http.Request) (string, error) {
	jwtToken := request.Header.Get("Authorization")
	if jwtToken == "" {
		err := map[string]interface{}{"ValidateAuthorization": fmt.Errorf("missing authorization Bearer").Error()}
		s.HandleRouteError(writer, err, http.StatusUnauthorized)
		return "", nil
	}
	return jwtToken[len("Bearer "):], nil
}
