package business_logic

import "moto-management-server/business_logic/models"

func (b *BusinessLogic) UpdateUser(user models.User) (models.User, error) {

	newUser, newUserErr := b.mongoClient.UpdateUser(fromBlUserToMongoUser(user))
	if newUserErr != nil {
		return models.User{}, newUserErr
	}

	return fromMongoUserToBlUser(newUser), nil
}
