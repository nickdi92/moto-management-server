package business_logic

import (
	"errors"
	"moto-management-server/business_logic/models"
)

func (b *BusinessLogic) GetUserByUsername(username string) (models.User, error) {
	if username == "" {
		return models.User{}, errors.New("no username provided")
	}
	mongoUser, mongoFindErr := b.mongoClient.GetUserByUsername(username)

	if mongoFindErr != nil {
		return models.User{}, errors.New(mongoFindErr.Error())
	}

	return fromMongoUserToBlUser(mongoUser), nil
}
