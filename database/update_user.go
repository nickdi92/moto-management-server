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
	usersCollections := m.mongoClient.Database(os.Getenv("MONGODB_DATABASE")).Collection(os.Getenv("MONGODB_USERS_COLLECTIONS"))

	filter := bson.D{{"_id", userToUpdate.ID}}
	//updateSingleFields := make([]bson.M, 0)
	update := bson.M{"$set": bson.M{"motorcycles": userToUpdate.Motorcycles}}

	//if userToUpdate.Name != "" {
	//	updateSingleFields = append(updateSingleFields, bson.M{
	//		"name": userToUpdate.Name,
	//	})
	//	update["$set"] = append(update["$set"], bson.M{
	//		"name": userToUpdate.Name,
	//	})
	//}
	//
	//if userToUpdate.Lastname != "" {
	//	updateSingleFields = append(updateSingleFields, bson.M{
	//		"lastname": userToUpdate.Lastname,
	//	})
	//}
	//
	//if userToUpdate.Token != "" {
	//	updateSingleFields = append(updateSingleFields, bson.M{
	//		"token": userToUpdate.Token,
	//	})
	//	updateSingleFields = append(updateSingleFields, bson.M{
	//		"expire_at": userToUpdate.ExpireAt,
	//	})
	//}
	//
	//if userToUpdate.Lastname != "" {
	//	updateSingleFields = append(updateSingleFields, bson.M{
	//		"lastname": userToUpdate.Lastname,
	//	})
	//}
	//
	//if userToUpdate.Motorcycles != nil {
	//	updateSingleFields = append(updateSingleFields, bson.M{
	//		"motorcycles": userToUpdate.Motorcycles,
	//	})
	//}
	//
	//updateSingleFields = append(updateSingleFields, bson.M{
	//	"updated_at": time.Now(),
	//})
	//
	//updateSingleFields = append(updateSingleFields, bson.M{
	//	"is_logged_in": userToUpdate.IsLoggedIn,
	//})
	//
	//
	//jsonMotors, _ := json.Marshal(userToUpdate.Motorcycles)
	//fmt.Printf("jsonMotors: %s", jsonMotors)

	_, err := usersCollections.UpdateOne(ctx, filter, update)

	if err != nil {
		return mongoUser, err
	}

	mongoUser, _ = m.GetUserByUsername(userToUpdate.Username)
	return mongoUser, nil
}
