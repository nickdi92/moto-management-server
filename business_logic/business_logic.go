package business_logic

import (
	"moto-management-server/business_logic/models"
	"moto-management-server/database"
	"moto-management-server/utils"
	"os"
)

type BusinessLogicInterface interface {
	NewBusinessLogic() *BusinessLogic

	GetUserByUsername(username string) (models.User, error)
	CreateNewUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)

	DeleteMotorbike(username string, licensePlate string) (bool, error)
	GetMotorcycleByLicensePlate(username string, licensePlate string) (models.Motorcycle, error)
	AddFuelToMotorcycle(username string, licensePlate string, fuel models.FuelSupplies) (models.User, error)

	// Services
	AddServiceToMotorcycle(username string, licensePlate string, service models.Service) (models.User, error)
	RemoveServiceFromMotorcycle(username string, licensePlate string, serviceId string) (models.User, error)
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
