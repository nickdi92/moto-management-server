package business_logic

import (
	"moto-management-server/business_logic/models"
)

func (b *BusinessLogic) UpdateUser(user models.User) (models.User, error) {

	blUser, getUserErr := b.GetUserByUsername(user.Username)
	if getUserErr != nil {
		return models.User{}, getUserErr
	}
	user.ID = blUser.ID
	newUser, newUserErr := b.mongoClient.UpdateUser(fromBlUserToMongoUser(user))
	if newUserErr != nil {
		return models.User{}, newUserErr
	}

	return fromMongoUserToBlUser(newUser), nil
}
