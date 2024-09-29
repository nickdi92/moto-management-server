package business_logic

import (
	"errors"
	"moto-management-server/business_logic/models"
)

func (b *BusinessLogic) AddServiceToMotorcycle(username string, licensePlate string, service models.Service) (models.User, error) {
	if username == "" {
		return models.User{}, errors.New("no username provided")
	}
	blUSer, findErr := b.GetUserByUsername(username)
	if findErr != nil {
		return models.User{}, findErr
	}

	motorcycleToUpdate, getMotorcycleToUpdateErr := b.GetMotorcycleByLicensePlate(username, licensePlate)
	if getMotorcycleToUpdateErr != nil {
		return models.User{}, getMotorcycleToUpdateErr
	}

	services := make([]models.Service, 0)
	services = append(services, service)
	canUpdate := false

	if motorcycleToUpdate.Service != nil {
		services = append(services, motorcycleToUpdate.Service...)
	}

	motorcycleToUpdate.Service = services
	for index, mt := range blUSer.Motorcycles {
		if mt.ID == motorcycleToUpdate.ID {
			blUSer.Motorcycles[index].Service = motorcycleToUpdate.Service
			canUpdate = true
			break
		}
	}
	if canUpdate {
		updatedUser, updateErr := b.UpdateUser(blUSer)
		if updateErr != nil {
			return models.User{}, updateErr
		}
		return updatedUser, nil
	}

	return models.User{}, nil
}
