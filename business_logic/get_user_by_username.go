package business_logic

import (
	"errors"
)

func (b *BusinessLogic) GetUserByUsername(username string) (User, error) {
	user := User{}
	if username == "" {
		return user, errors.New("No username provided") //@TODO improve errors
	}
	// mongo, _ := b.mongoClient.NewMongoClient() @TODO This is wrong
	_, mongoFindErr := b.mongoClient.GetUserByUsername(username)

	if mongoFindErr != nil {
		return user, errors.New(mongoFindErr.Error())
	}
	return user, nil
}
