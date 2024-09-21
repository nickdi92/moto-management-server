package server

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func (s *MotoManagementServer) ValidateRequest(structToValidate interface{}) error {
	validate := validator.New()
	validationErr := validate.Struct(structToValidate)
	if validationErr != nil {
		return validationErr
	}
	return nil
}

func (s *MotoManagementServer) ValidateJwtToken(incomingJwtToken string, savedJwtToken string) error {
	token, err := jwt.Parse(incomingJwtToken, func(token *jwt.Token) (interface{}, error) {
		return savedJwtToken, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("error validating JWT")
	}

	return nil
}

func (s *MotoManagementServer) ValidateAuthorization(writer http.ResponseWriter, request *http.Request) (string, error) {
	jwtToken := request.Header.Get("Authorization")
	if jwtToken == "" {
		err := map[string]interface{}{"loginRouteErr": fmt.Errorf("missing authorization Bearer").Error()}
		writer.WriteHeader(http.StatusUnauthorized)
		s.HandleRouteError(writer, err)
		return "", nil
	}
	return jwtToken, nil
}
