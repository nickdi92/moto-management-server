package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"moto-management-server/errors"
	"time"
)

func (m *MotoManagementMongoClient) GetUserByUsername(username string) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var user User
	findErr := m.usersCollections.FindOne(ctx, bson.D{{"username", username}}).Decode(&user)
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
