package server

import (
	"fmt"
	"moto-management-server/business_logic"
	"moto-management-server/utils"
	"net/http"
)

func (s *MotoManagementServer) NewMotoManagementServer() (*MotoManagementServer, error) {
	bl := business_logic.BusinessLogic{}
	bl.NewBusinessLogic()

	server := &MotoManagementServer{
		Addr:          ":8080",
		businessLogic: bl,
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
