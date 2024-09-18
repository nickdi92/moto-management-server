package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"moto-management-server/errors"
	"os"
	"time"
)

func (m *MotoManagementMongoClient) GetUserByUsername(username string) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if m == nil {
		fmt.Println("CANDISCO")
	} else {
		fmt.Println("ZIOMANESKIN")
	}

	var user User
	usersCollections := m.mongoClient.
		Database(os.Getenv("MONGODB_DATABASE")).
		Collection(os.Getenv("MONGODB_USERS_COLLECTIONS"))

	sr := usersCollections.FindOne(ctx, bson.D{{"username", username}})
	if sr == nil {
		return User{}, errors.MongoErrors{
			Code:    errors.MongoErrorCode_NoDocumentsFound,
			Message: "No users found",
		}
	}
	findErr := sr.Decode(&user)
	if findErr != nil {
		if findErr == mongo.ErrNoDocuments {
			return User{}, errors.MongoErrors{
				Code:    errors.MongoErrorCode_NoDocumentsFound,
				Message: findErr.Error(),
			}
		}
		return User{}, errors.MongoErrors{
			Code:    errors.MongoErrorCode_ErrorOnSearchingIntoDatabase,
			Message: findErr.Error(),
		}
	}
	return user, nil
}
