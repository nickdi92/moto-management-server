package business_logic

import (
	"errors"
	"fmt"
	"moto-management-server/business_logic/models"
)

func (b *BusinessLogic) CreateNewUser(userToCreate models.User) (models.User, error) {
	prevUser, _ := b.mongoClient.GetUserByUsername(userToCreate.Username)
	blPrevUser := fromMongoUserToBlUser(prevUser)
	if blPrevUser.ID != "" {
		return models.User{}, errors.New(fmt.Sprintf("user %s already exist", userToCreate.Username))
	}

	newUser, newUserErr := b.mongoClient.CreateNewUser(fromBlUserToMongoUser(userToCreate))
	if newUserErr != nil {
		return models.User{}, newUserErr
	}

	return fromMongoUserToBlUser(newUser), nil
}
