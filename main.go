package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	server2 "moto-management-server/server"
	"moto-management-server/utils"
	"os"
)

func main() {

	/***********************************
				.ENV
	***********************************/
	err := godotenv.Load()
	if err != nil {
		utils.ErrorOutput("Error loading .env file")
		log.Fatal("Error loading .env file")
	}

	utils.SuccessOutput("File .env loaded successfully !")

	//businessLogic := business_logic.BusinessLogic{}
	//businessLogic.NewBusinessLogic()

	/***********************************
				WEBSERVER
	***********************************/
	motoServer := server2.MotoManagementServer{}
	_, serverError := motoServer.NewMotoManagementServer()
	if serverError != nil {
		utils.ErrorOutput(fmt.Sprintf("error is: %v", err))
		os.Exit(1)
	}
}
