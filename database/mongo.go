package database

import (
	"context"
	"moto-management-server/database/models"
	"moto-management-server/errors"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MotoManagementMongoClient struct {
	mongoClient *mongo.Client
}

type MotoManagementMongoClientInterface interface {
	NewMongoClient() (*MotoManagementMongoClient, error)

	GetUserByUsername(username string) (models.User, error)
	CreateNewUser(userToCreate models.User) (models.User, error)
	UpdateUser(userToUpdate models.User) (models.User, error)

	DeleteMotorbike(username string, licensePlate string) (bool, error)
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

	// Check if MongoDB connection was successful
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, errors.MongoErrors{
			Code:    errors.MongoErrorCode_FailedToPing,
			Message: err.Error(),
		}
	}

	usersCollections := client.
		Database(os.Getenv("MONGODB_DATABASE")).
		Collection(os.Getenv("MONGODB_USERS_COLLECTIONS"))

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
		mongoClient: client,
	}, nil
}
