package database

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (m *MotoManagementMongoClient) UpdateUser(userToUpdate User) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	mongoUser := User{}
	userToUpdate.UpdatedAt = time.Now()

	uusersCollections := m.mongoClient.Database(os.Getenv("MONGODB_DATABASE")).Collection(os.Getenv("MONGODB_USERS_COLLECTIONS"))

	// @todo Add all fields to update
	filter := bson.D{{"_id", userToUpdate.ID}}
	update := bson.D{{"$set", bson.D{
		{
			"is_logged_in", userToUpdate.IsLoggedIn,
		},
	}}}

	_, err := uusersCollections.UpdateOne(ctx, filter, update)

	if err != nil {
		return mongoUser, err
	}

	mongoUser, _ = m.GetUserByUsername(userToUpdate.Username)
	return mongoUser, nil
}
