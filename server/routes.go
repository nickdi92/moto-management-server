package server

import (
	"fmt"
	"moto-management-server/errors"
	"moto-management-server/server/routes"
	"net/http"
)

const webPrefix = "web"

func (s *MotoManagementServer) RegisterRoutes() {
	s.routes = make(Routes)

	// Add manual routes
	privateRoutes := make(Routes)
	privateRoutes["auth"] = routes.AuthRoute
	//privateRoutes["login"] = routes.LoginRoute

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
		http.HandleFunc(url, routeHandler)
	}
	return nil
}
