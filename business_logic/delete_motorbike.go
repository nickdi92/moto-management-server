package business_logic

func (b *BusinessLogic) DeleteMotorbike(username string, licensePlate string) (bool, error) {
	_, getUserErr := b.GetUserByUsername(username)
	if getUserErr != nil {
		return false, getUserErr
	}

	isDeleted, deleteErr := b.mongoClient.DeleteMotorbike(username, licensePlate)
	if deleteErr != nil {
		return false, deleteErr
	}
	return isDeleted, nil
}
