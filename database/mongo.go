package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"moto-management-server/errors"
	"os"
	"time"
)

type MotoManagementMongoClient struct {
	mongoClient      *mongo.Client
	usersCollections *mongo.Collection
}

type MotoManagementMongoClientInterface interface {
	NewMongoClient() (*MotoManagementMongoClient, error)
	GetUserByUsername(username string) (User, error)
}

func (m *MotoManagementMongoClient) NewMongoClient() (*MotoManagementMongoClient, error) {

	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, errors.MongoErrors{
			Code:    errors.MongoErrorCode_FailedToConnect,
			Message: err.Error(),
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// Check if MongoDB connection was successful
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, errors.MongoErrors{
			Code:    errors.MongoErrorCode_FailedToPing,
			Message: err.Error(),
		}
	}

	usersCollections := client.Database(os.Getenv("MONGODB_DATABASE")).Collection(os.Getenv("MONGODB_USERS_COLLECTIONS"))

	/** Create Collections indexes */
	usersIndexes := make([]mongo.IndexModel, 0)
	usersIndexes = append(usersIndexes, mongo.IndexModel{
		Keys: bson.D{
			{
				"id",
				-1,
			},
		},
		Options: options.Index().SetUnique(true),
	})

	usersIndexes = append(usersIndexes, mongo.IndexModel{
		Keys: bson.D{

			{
				"username",
				1,
			},
		},
		Options: options.Index().SetUnique(true),
	})

	usersIndexes = append(usersIndexes, mongo.IndexModel{
		Keys: bson.D{
			{
				"motorcycles.license_plate",
				1,
			},
		},
		Options: options.Index().SetUnique(true),
	})
	_, indexErr := usersCollections.Indexes().CreateMany(ctx, usersIndexes)
	if indexErr != nil {
		return nil, errors.MongoErrors{
			Code:    errors.MongoErrorCode_CreateIndexError,
			Message: err.Error(),
		}
	}
	/******* end ******/

	return &MotoManagementMongoClient{
		mongoClient:      client,
		usersCollections: usersCollections,
	}, nil
}
