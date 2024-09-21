package business_logic

import (
	"moto-management-server/database"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func fromMongoUserToBlUser(mongoUser database.User) User {
	id := mongoUser.ID.Hex()
	if mongoUser.ID.IsZero() {
		id = ""
	}
	return User{
		ID:         id,
		Username:   mongoUser.Username,
		Name:       mongoUser.Name,
		Lastname:   mongoUser.Lastname,
		Password:   mongoUser.Password,
		Email:      mongoUser.Email,
		Token:      mongoUser.Token,
		ExpireAt:   &mongoUser.ExpireAt,
		CreatedAt:  &mongoUser.CreatedAt,
		UpdatedAt:  &mongoUser.UpdatedAt,
		IsLoggedIn: mongoUser.IsLoggedIn,
	}
}

func fromBlUserToMongoUser(blUser User) database.User {
	mongoUser := database.User{
		Username:   blUser.Username,
		Name:       blUser.Name,
		Lastname:   blUser.Lastname,
		Password:   blUser.Password,
		Email:      blUser.Email,
		Token:      blUser.Token,
		ExpireAt:   *blUser.ExpireAt,
		IsLoggedIn: blUser.IsLoggedIn,
	}

	if blUser.ID != "" {
		mongoUser.ID, _ = primitive.ObjectIDFromHex(blUser.ID)
	}
	return mongoUser
}
