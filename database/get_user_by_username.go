package database

import (
	"context"
	"moto-management-server/database/models"
	"moto-management-server/errors"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *MotoManagementMongoClient) GetUserByUsername(username string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var user models.User
	usersCollections := m.mongoClient.
		Database(os.Getenv("MONGODB_DATABASE")).
		Collection(os.Getenv("MONGODB_USERS_COLLECTIONS"))

	sr := usersCollections.FindOne(ctx, bson.D{{"username", username}})
	if sr == nil {
		return models.User{}, errors.MongoErrors{
			Code:    errors.MongoErrorCode_NoDocumentsFound,
			Message: "No users found",
		}
	}

	findErr := sr.Decode(&user)
	if findErr != nil {
		if findErr == mongo.ErrNoDocuments {
			return models.User{}, errors.MongoErrors{
				Code:    errors.MongoErrorCode_NoDocumentsFound,
				Message: findErr.Error(),
			}
		}
		return models.User{}, errors.MongoErrors{
			Code:    errors.MongoErrorCode_ErrorOnSearchingIntoDatabase,
			Message: findErr.Error(),
		}
	}
	return user, nil
}
