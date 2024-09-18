package server

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

func (s *MotoManagementServer) ValidateRequest(request *http.Request, rules govalidator.MapData) error {
	opts := govalidator.Options{
		Request: request,
		Rules:   rules,
	}

	v := govalidator.New(opts)
	validationErr := v.Validate()
	if len(validationErr) > 0 {
		err := errors.New(fmt.Sprintf("Error validating request: %v", validationErr.Encode()))
		return err
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
