package business_logic

import (
	"errors"
	"fmt"
)

func (b *BusinessLogic) CreateNewUser(userToCreate User) (User, error) {
	prevUser, _ := b.mongoClient.GetUserByUsername(userToCreate.Username)
	blPrevUser := fromMongoUserToBlUser(prevUser)
	if blPrevUser.ID != "" {
		return User{}, errors.New(fmt.Sprintf("user %s already exist", userToCreate.Username))
	}

	newUser, newUserErr := b.mongoClient.CreateNewUser(fromBlUserToMongoUser(userToCreate))
	if newUserErr != nil {
		return User{}, newUserErr
	}

	return fromMongoUserToBlUser(newUser), nil
}