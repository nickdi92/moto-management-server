package business_logic

import (
	"moto-management-server/database"
	"moto-management-server/utils"
	"os"
)

type BusinessLogicInterface interface {
	NewBusinessLogic() *BusinessLogic
	GetUserByUsername(username string) (User, error)
}

func (b *BusinessLogic) NewBusinessLogic() *BusinessLogic {
	mongoCl := database.MotoManagementMongoClient{}
	client, mongoErr := mongoCl.NewMongoClient()
	if mongoErr != nil {
		utils.ErrorOutput(mongoErr.Error())
		os.Exit(1)
	}
	utils.SuccessOutput("Connected to MongoDB ! ")
	return &BusinessLogic{
		mongoClient: client,
	}
}
