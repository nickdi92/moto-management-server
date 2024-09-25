package server

import (
	"encoding/json"
	"fmt"
	"moto-management-server/errors"
	"net/http"
)

const adminPrefix = "admin"

func (s *MotoManagementServer) RegisterRoutes() {
	s.routes = make(Routes)

	// Add manual routes
	privateRoutes := make(Routes)
	privateRoutes["token"] = TokenRoute
	privateRoutes["register"] = RegisterRoute
	privateRoutes["login"] = LoginRoute

	/*------------------------------------------------*
	 *					 USER ROUTES				  *
	 *------------------------------------------------*/

	privateRoutes["user/get"] = GetUserRoute

	/*------------------------------------------------*
	 *				MOTORCYCLE ROUTES				  *
	 *------------------------------------------------*/

	privateRoutes["motorcycle/add"] = MotorcyclesAddRoute
	privateRoutes["motorcycle/delete"] = MotorcyclesDeleteRoute

	for url, routeHandler := range privateRoutes {
		// Building url like /web/auth, /web/login
		routeUrl := fmt.Sprintf("/%s/%s", adminPrefix, url)
		s.routes[routeUrl] = routeHandler
	}
}

func (s *MotoManagementServer) HandleRoutes() error {
	if s.routes == nil {
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

func (s *MotoManagementServer) HandleRouteError(writer http.ResponseWriter, err interface{}, status int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	_ = json.NewEncoder(writer).Encode(err)
}

func (s *MotoManagementServer) HandleResponse(writer http.ResponseWriter, result interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(result)
}
