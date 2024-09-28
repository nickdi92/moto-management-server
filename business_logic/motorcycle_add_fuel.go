package business_logic

import (
	"errors"
	"moto-management-server/business_logic/models"
)

func (b *BusinessLogic) AddFuelToMotorcycle(username string, licensePlate string, fuel models.FuelSupplies) (models.Motorcycle, error) {
	if username == "" {
		return models.Motorcycle{}, errors.New("no username provided")
	}
	blUSer, findErr := b.GetUserByUsername(username)
	if findErr != nil {
		return models.Motorcycle{}, findErr
	}

	motorcycleToUpdate, getMotorcycleToUpdateErr := b.GetMotorcycleByLicensePlate(username, licensePlate)
	if getMotorcycleToUpdateErr != nil {
		return models.Motorcycle{}, getMotorcycleToUpdateErr
	}

	fuelSupplies := make([]models.FuelSupplies, 0)
	fuelSupplies = append(fuelSupplies, fuel)

	if motorcycleToUpdate.FuelSupplies != nil {

		fuelSupplies = append(fuelSupplies, motorcycleToUpdate.FuelSupplies...)
	}

	motorcycleToUpdate.FuelSupplies = fuelSupplies

	canUpdate := false
	for index, mt := range blUSer.Motorcycles {
		if mt.ID == motorcycleToUpdate.ID {
			blUSer.Motorcycles[index].FuelSupplies = motorcycleToUpdate.FuelSupplies
			canUpdate = true
			break
		}
	}

	if canUpdate {
		_, updateErr := b.UpdateUser(blUSer)
		if updateErr != nil {
			return models.Motorcycle{}, updateErr
		}
		return motorcycleToUpdate, nil
	}

	return models.Motorcycle{}, nil
}
