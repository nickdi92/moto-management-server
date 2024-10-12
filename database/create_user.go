package database

import (
	"context"
	"moto-management-server/database/models"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *MotoManagementMongoClient) CreateNewUser(userToCreate models.User) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	mongoUser := models.User{}
	userToCreate.CreatedAt = time.Now()
	userToCreate.UpdatedAt = time.Now()

	usersCollections := m.mongoClient.
		Database(os.Getenv("MONGODB_DATABASE")).
		Collection(os.Getenv("MONGODB_USERS_COLLECTIONS"))

	userToCreate.ID = primitive.NewObjectID()
	_, err := usersCollections.InsertOne(ctx, userToCreate)

	if err != nil {
		return mongoUser, err
	}

	mongoUser, _ = m.GetUserByUsername(userToCreate.Username)
	return mongoUser, nil
}
