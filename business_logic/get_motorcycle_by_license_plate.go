package business_logic

import (
	"errors"
	"moto-management-server/business_logic/models"
)

func (b *BusinessLogic) GetMotorcycleByLicensePlate(username string, licensePlate string) (models.Motorcycle, error) {
	if username == "" {
		return models.Motorcycle{}, errors.New("no username provided")
	}
	blUSer, findErr := b.GetUserByUsername(username)
	if findErr != nil {
		return models.Motorcycle{}, findErr
	}

	var moto models.Motorcycle

	if blUSer.Motorcycles == nil {
		return moto, errors.New("no motorcycles found for given username")
	}

	for _, mt := range blUSer.Motorcycles {
		if mt.LicensePlate == licensePlate {
			moto = mt
			break
		}
	}
	return moto, nil
}
