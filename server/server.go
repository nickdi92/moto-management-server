package server

import (
	"log"
	"net/http"
)

func NewMotoManagementServer() (*MotoManagementServer, error) {
	server := &MotoManagementServer{
		Addr: ":8080",
	}

	err := server.RegisterRoutes()
	if err != nil {
		return nil, err
	}

	err = server.HandleRoutes()
	if err != nil {
		return nil, err
	}

	log.Fatal(http.ListenAndServe(server.Addr, nil))

	return server, nil
}
