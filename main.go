package main

import (
	server2 "moto-management-server/server"
)

func main() {

	db := client.Database("moto_management")
	db.Collection("users")

	//-----------------------------------

	// Create a new HTTP server
	_, _ = server2.NewMotoManagementServer()
}
