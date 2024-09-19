package business_logic

import (
	"errors"
)

func (b *BusinessLogic) GetUserByUsername(username string) (User, error) {
	if username == "" {
		return User{}, errors.New("No username provided") //@TODO improve errors
	}
	// mongo, _ := b.mongoClient.NewMongoClient() @TODO This is wrong
	mongoUser, mongoFindErr := b.mongoClient.GetUserByUsername(username)

	if mongoFindErr != nil {
		return User{}, errors.New(mongoFindErr.Error())
	}
	
	return fromMongoUserToBlUser(mongoUser), nil
}
