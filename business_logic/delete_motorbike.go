package business_logic

import "moto-management-server/business_logic/models"

func (b *BusinessLogic) DeleteMotorbike(username string, licensePlate string) (bool, error) {
	gotUser, getUserErr := b.GetUserByUsername(username)
	if getUserErr != nil {
		return false, getUserErr
	}

	if gotUser.Motorcycles != nil && len(gotUser.Motorcycles) > 0 {
		motorCycles := make([]models.Motorcycle, 0)
		for _, mt := range gotUser.Motorcycles {
			// remove the motorbike which licensePlate is equal to the given license_plate
			if mt.LicensePlate != licensePlate {
				motorCycles = append(motorCycles, mt)
			}
		}
		gotUser.Motorcycles = motorCycles
	}

	_, updateErr := b.UpdateUser(gotUser)
	if updateErr != nil {
		return false, updateErr
	}
	return true, nil
}
