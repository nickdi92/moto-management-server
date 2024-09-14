package business_logic

import (
	"moto-management-server/database"
)

type BusinessLogic struct {
	mongoClient *database.MotoManagementMongoClient
}
