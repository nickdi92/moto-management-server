package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"time"
)

func (m *MotoManagementMongoClient) DeleteMotorbike(username string, licensePlate string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	usersCollections := m.mongoClient.Database(os.Getenv("MONGODB_DATABASE")).Collection(os.Getenv("MONGODB_USERS_COLLECTIONS"))
	filter := bson.D{{"username", username}, {"motorcycles.license_plate", licensePlate}}

	r, err := usersCollections.DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return r.DeletedCount > 0, nil
}
