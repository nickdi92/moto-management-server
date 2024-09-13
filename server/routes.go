package server

import (
	"fmt"
	"moto-management-server/errors"
	"moto-management-server/server/routes"
	"net/http"
)

const webPrefix = "web"

func (s *MotoManagementServer) RegisterRoutes() error {
	s.routes = make(Routes)

	// Add manual routes
	privateRoutes := make(Routes)
	privateRoutes["auth"] = routes.AuthRoute
	privateRoutes["login"] = routes.LoginRoute

	/*-------------------------------------------------*/
	for url, routeHandler := range s.routes {
		// Building url like /web/auth, /web/login
		routeUrl := fmt.Sprintf("/%s/%s", webPrefix, url)
		s.routes[routeUrl] = routeHandler
	}

	return nil
}

func (s *MotoManagementServer) HandleRoutes() error {
	if s.routes == nil {
		return errors.RouteErrors{
			Code:    errors.RouteErrorCode_NoRoutesRegistered,
			Message: "No routes registered",
		}
	}
	for url, route := range s.routes {
		http.HandleFunc(url, route)
	}
	return nil
}
