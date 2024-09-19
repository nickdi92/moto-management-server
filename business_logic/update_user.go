package business_logic

func (b *BusinessLogic) UpdateUser(user User) (User, error) {

	newUser, newUserErr := b.mongoClient.UpdateUser(fromBlUserToMongoUser(user))
	if newUserErr != nil {
		return User{}, newUserErr
	}

	return fromMongoUserToBlUser(newUser), nil
}