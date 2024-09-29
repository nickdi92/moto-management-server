package database

import (
	"context"
	"moto-management-server/database/models"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *MotoManagementMongoClient) UpdateUser(userToUpdate models.User) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	mongoUser := models.User{}
	usersCollections := m.mongoClient.Database(os.Getenv("MONGODB_DATABASE")).Collection(os.Getenv("MONGODB_USERS_COLLECTIONS"))

	filter := bson.D{{"_id", userToUpdate.ID}}
	updateSingleFields := bson.M{}

	if userToUpdate.Name != "" {
		updateSingleFields["name"] = userToUpdate.Name
	}

	if userToUpdate.Lastname != "" {
		updateSingleFields["lastname"] = userToUpdate.Lastname
	}

	if userToUpdate.Token != "" {
		updateSingleFields["token"] = userToUpdate.Token
		updateSingleFields["expire_at"] = userToUpdate.ExpireAt
	}

	if userToUpdate.Motorcycles != nil {
		for index, mt := range userToUpdate.Motorcycles {
			if mt.ID.IsZero() {
				userToUpdate.Motorcycles[index].ID = primitive.NewObjectID()
			}
			for fuelIndex, f := range mt.FuelSupplies {
				if f.ID.IsZero() {
					userToUpdate.Motorcycles[index].FuelSupplies[fuelIndex].ID = primitive.NewObjectID()
				}
			}

			for serviceIndex, service := range mt.Service {
				if service.ID.IsZero() {
					userToUpdate.Motorcycles[index].Service[serviceIndex].ID = primitive.NewObjectID()
				}
			}
		}
		updateSingleFields["motorcycles"] = userToUpdate.Motorcycles
	}

	updateSingleFields["updated_at"] = time.Now()
	updateSingleFields["is_logged_in"] = userToUpdate.IsLoggedIn

	update := bson.M{"$set": updateSingleFields}
	_, err := usersCollections.UpdateOne(ctx, filter, update)

	if err != nil {
		return mongoUser, err
	}

	mongoUser, _ = m.GetUserByUsername(userToUpdate.Username)
	return mongoUser, nil
}
