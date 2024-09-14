package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"moto-management-server/errors"
	"os"
)

func NewMongoClient() error {

	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return errors.MongoErrors{
			Code:    errors.MongoErrorCode_FailedToConnect,
			Message: err.Error(),
		}
	}

	defer client.Disconnect(context.Background())

	// Check if MongoDB connection was successful
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return errors.MongoErrors{
			Code:    errors.MongoErrorCode_FailedToPing,
			Message: err.Error(),
		}
	}
	return nil
}
