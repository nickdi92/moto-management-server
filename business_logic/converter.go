package business_logic

import "moto-management-server/database"

func fromMongoUserToBlUser(mongoUser database.User) User {
	id := mongoUser.ID.Hex()
	if mongoUser.ID.IsZero() {
		id = ""
	}
	return User{
		ID:        id,
		Username:  mongoUser.Username,
		Name:      mongoUser.Name,
		Lastname:  mongoUser.Lastname,
		Password:  mongoUser.Password,
		Token:     mongoUser.Token,
		ExpireAt:  &mongoUser.ExpireAt,
		CreatedAt: &mongoUser.CreatedAt,
		UpdatedAt: &mongoUser.UpdatedAt,
	}
}

func fromBlUserToMongoUser(blUser User) database.User {
	return database.User{
		Username: blUser.Username,
		Name:     blUser.Name,
		Lastname: blUser.Lastname,
		Password: blUser.Password,
		Token:    blUser.Token,
		ExpireAt: *blUser.ExpireAt,
	}
}
