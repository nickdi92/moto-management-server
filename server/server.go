package server

import (
	"fmt"
	"log"
	"moto-management-server/business_logic"
	"moto-management-server/utils"
	"net/http"
)

func (s *MotoManagementServer) NewMotoManagementServer() (*MotoManagementServer, error) {
	bl := &business_logic.BusinessLogic{}
	server := &MotoManagementServer{
		Addr:          ":8080",
		businessLogic: bl.NewBusinessLogic(),
	}

	server.RegisterRoutes()

	err := server.HandleRoutes()
	if err != nil {
		return nil, err
	}

	utils.SuccessOutput("Webserver started")
	utils.InfoOutput(fmt.Sprintf("Listening on localhost%s", server.Addr))

	err = http.ListenAndServe(server.Addr, logRequest(http.DefaultServeMux))
	return server, err
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
