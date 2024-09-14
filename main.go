package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"moto-management-server/database"
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

	/***********************************
				MONGO
	***********************************/
	mongoErr := database.NewMongoClient()
	if mongoErr != nil {
		utils.ErrorOutput(mongoErr.Error())
		os.Exit(1)
	}
	utils.SuccessOutput("Connected to MongoDB !")

	/***********************************
				WEBSERVER
	***********************************/
	_, serverError := server2.NewMotoManagementServer()
	if serverError != nil {
		utils.ErrorOutput(fmt.Sprintf("error is: %v", err))
		os.Exit(1)
	}
}
