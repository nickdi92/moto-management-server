package server

import (
	"moto-management-server/business_logic"
	"net/http"
)

type Route func(writer http.ResponseWriter, request *http.Request)
type RouteHandler func(b business_logic.BusinessLogic, url string, route Route)
type Routes map[string]Route

type MotoManagementServer struct {
	Addr          string // Listen port. Normally :8080
	routes        Routes // Need to register route handlers
	businessLogic business_logic.BusinessLogic
}

type MotoManagementServerInterface interface {
	NewMotoManagementServer() (*MotoManagementServer, error)
	RegisterRoutes()
	HandleRoutes() error
}
