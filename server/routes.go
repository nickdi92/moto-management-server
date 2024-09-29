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

	/*------------------------------------------------*
	 *					 USER ROUTES				  *
	 *------------------------------------------------*/

	privateRoutes["user/refresh-token"] = RefreshTokenRoute
	privateRoutes["user/create"] = UserCreateRoute
	privateRoutes["user/login"] = LoginRoute
	privateRoutes["user/get"] = GetUserRoute
	privateRoutes["user/update"] = UserUpdateRoute

	/*------------------------------------------------*
	 *				MOTORCYCLE ROUTES				  *
	 *------------------------------------------------*/

	privateRoutes["motorcycle/add"] = MotorcyclesAddRoute
	privateRoutes["motorcycle/delete"] = MotorcyclesDeleteRoute
	privateRoutes["motorcycle/getByLicensePlate"] = MotorcyclesGetByLicensePlateRoute

	/*------------------------------------------------*
	 *				  	FUEL ROUTES				   	  *
	 *------------------------------------------------*/

	privateRoutes["motorcycle/fuel/add"] = MotorcyclesAddFuelRoute

	/*------------------------------------------------*
	 *				  SERVICES ROUTES			   	  *
	 *------------------------------------------------*/

	privateRoutes["motorcycle/services/add"] = MotorcyclesAddServiceRoute
	privateRoutes["motorcycle/services/remove"] = MotorcyclesDeleteServiceRoute

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
