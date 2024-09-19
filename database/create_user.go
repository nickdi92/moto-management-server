package database

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *MotoManagementMongoClient) CreateNewUser(userToCreate User) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	mongoUser := User{}
	userToCreate.CreatedAt = time.Now()
	userToCreate.UpdatedAt = time.Now()

	uusersCollections := m.mongoClient.
		Database(os.Getenv("MONGODB_DATABASE")).
		Collection(os.Getenv("MONGODB_USERS_COLLECTIONS"))

	userToCreate.ID = primitive.NewObjectID()
	_, err := uusersCollections.InsertOne(ctx, userToCreate)

	if err != nil {
		return mongoUser, err
	}

	mongoUser, _ = m.GetUserByUsername(userToCreate.Username)
	return mongoUser, nil
}
