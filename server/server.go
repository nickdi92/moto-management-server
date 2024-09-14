package server

import (
	"fmt"
	"moto-management-server/utils"
	"net/http"
)

func NewMotoManagementServer() (*MotoManagementServer, error) {
	server := &MotoManagementServer{
		Addr: ":8080",
	}

	server.RegisterRoutes()

	err := server.HandleRoutes()
	if err != nil {
		return nil, err
	}

	utils.SuccessOutput("Webserver started")
	utils.InfoOutput(fmt.Sprintf("Listening on localhost%s", server.Addr))

	err = http.ListenAndServe(server.Addr, nil)
	return server, err
}
