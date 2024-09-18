package business_logic

import (
	"errors"
	"fmt"
)

func (b *BusinessLogic) GetUserByUsername(username string) (User, error) {
	user := User{}
	if username == "" {
		return user, errors.New("No username provided") //@TODO improve errors
	}
	if b.mongoClient == nil {
		fmt.Println("sei un cannibale")
	}
	_, mongoFindErr := b.mongoClient.GetUserByUsername(username)

	if mongoFindErr != nil {
		return user, errors.New(mongoFindErr.Error())
	}
	return user, nil
}
