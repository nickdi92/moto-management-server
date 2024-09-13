package server

import "net/http"

type Route func(writer http.ResponseWriter, request *http.Request)
type Routes map[string]Route

type MotoManagementServer struct {
	Addr   string // Listen port. Normally :8080
	routes Routes // Need to register route handlers
}
