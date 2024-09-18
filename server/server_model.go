package server

import (
	"github.com/thedevsaddam/govalidator"
	"moto-management-server/business_logic"
	"net/http"
)

type Route func(s *MotoManagementServer, writer http.ResponseWriter, request *http.Request)
type Routes map[string]Route

type MotoManagementServer struct {
	Addr          string // Listen port. Normally :8080
	routes        Routes // Need to register route handlers
	businessLogic *business_logic.BusinessLogic
}

type MotoManagementServerInterface interface {
	NewMotoManagementServer() (*MotoManagementServer, error)
	RegisterRoutes()
	HandleRoutes() error

	// Validate request methods
	ValidateRequest(request *http.Request, rules govalidator.MapData) error
	ValidateJwtToken(incomingJwtToken string, savedJwtToken string) error
}
