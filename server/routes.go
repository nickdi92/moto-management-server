package server

import (
	"encoding/json"
	"fmt"
	"moto-management-server/errors"
	"net/http"
)

const webPrefix = "web"

func (s *MotoManagementServer) RegisterRoutes() {
	s.routes = make(Routes)

	// Add manual routes
	privateRoutes := make(Routes)
	privateRoutes["auth"] = AuthRoute
	privateRoutes["register"] = RegisterRoute
	privateRoutes["login"] = LoginRoute

	/*-------------------------------------------------*/
	for url, routeHandler := range privateRoutes {
		// Building url like /web/auth, /web/login
		routeUrl := fmt.Sprintf("/%s/%s", webPrefix, url)
		s.routes[routeUrl] = routeHandler
	}
}

func (s *MotoManagementServer) HandleRoutes() error {
	if s.routes == nil || len(s.routes) == 0 {
		return errors.RouteError{
			Code:    errors.RouteErrorCode_NoRoutesRegistered,
			Message: "No routes registered",
		}
	}

	for url, routeHandler := range s.routes {
		http.HandleFunc(url, func(writer http.ResponseWriter, request *http.Request) {
			routeHandler(s, writer, request)
		})
	}
	return nil
}

func (s *MotoManagementServer) HandleRouteError(writer http.ResponseWriter, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(err)
}

func (s *MotoManagementServer) HandleResponse(writer http.ResponseWriter, result interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(result)
	return
}
