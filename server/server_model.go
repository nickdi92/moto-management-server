package server

import (
	"moto-management-server/business_logic"
	"moto-management-server/utils/token"
	"net/http"
)

type Route func(s *MotoManagementServer, writer http.ResponseWriter, request *http.Request)
type Routes map[string]Route

type MotoManagementServer struct {
	Addr          string // Listen port. Normally :8080
	routes        Routes // Need to register route handlers
	businessLogic *business_logic.BusinessLogic
	token         *token.Token
}

type Message struct {
	Status string `json:"status"`
	Info   string `json:"info"`
}

type MotoManagementServerInterface interface {
	NewMotoManagementServer() (*MotoManagementServer, error)
	RegisterRoutes()
	HandleRoutes() error

	// Validate request methods
	ValidateRequest(structToValidate interface{}) error
	ValidateAuthorization(writer http.ResponseWriter, request *http.Request) (string, error)
	ValidateJwtToken(incomingJwtToken string, savedJwtToken string) error
}
