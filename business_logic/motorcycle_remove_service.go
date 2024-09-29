package business_logic

import (
	"errors"
	"moto-management-server/business_logic/models"
)

func (b *BusinessLogic) RemoveServiceFromMotorcycle(username string, licensePlate string, serviceId string) (models.User, error) {
	if username == "" {
		return models.User{}, errors.New("no username provided")
	}
	blUSer, findErr := b.GetUserByUsername(username)
	if findErr != nil {
		return models.User{}, findErr
	}

	_, getMotorcycleToUpdateErr := b.GetMotorcycleByLicensePlate(username, licensePlate)
	if getMotorcycleToUpdateErr != nil {
		return models.User{}, getMotorcycleToUpdateErr
	}

	for motoIndex, mt := range blUSer.Motorcycles {
		services := make([]models.Service, 0)
		for _, service := range mt.Service {
			if service.ID != serviceId {
				services = append(services, service)
			}
		}
		blUSer.Motorcycles[motoIndex].Service = services
	}

	updatedUser, updateErr := b.UpdateUser(blUSer)
	if updateErr != nil {
		return models.User{}, updateErr
	}

	return updatedUser, nil
}
