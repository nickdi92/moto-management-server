package business_logic

import (
	"errors"
	"fmt"
	"github.com/r3labs/diff/v3"
	"moto-management-server/business_logic/models"
	"strings"
)

func (b *BusinessLogic) UpdateServiceToMotorcycle(username string, licensePlate string, service models.Service) (models.User, error) {
	if username == "" {
		return models.User{}, errors.New("no username provided")
	}

	if service.ID == "" {
		return models.User{}, errors.New("serviceId is a required field")
	}

	blUser, findErr := b.GetUserByUsername(username)
	if findErr != nil {
		return models.User{}, findErr
	}

	motorcycleToUpdate, getMotorcycleToUpdateErr := b.GetMotorcycleByLicensePlate(username, licensePlate)
	if getMotorcycleToUpdateErr != nil {
		return models.User{}, getMotorcycleToUpdateErr
	}

	for index, mt := range motorcycleToUpdate.Service {
		if mt.ID != service.ID {
			continue
		}

		differ, _ := diff.NewDiffer(diff.TagName("json"))
		changeLog, changelogErr := differ.Diff(mt, service)
		if changelogErr != nil {
			return models.User{}, changelogErr
		}
		patchLog := differ.Patch(changeLog, &mt)
		if patchLog.ErrorCount() > 0 && patchLog.HasErrors() {
			for _, p := range patchLog {
				if p.Errors != nil {
					errorString := fmt.Sprintf(
						"Error on field: %s. From value %v. To Value %v. Final Error is: %s",
						strings.Join(p.Path, "->"), p.From, p.To, p.Errors.Error(),
					)
					return models.User{}, errors.New(errorString)
				}
			}
		}
		if patchLog.Applied() {
			motorcycleToUpdate.Service[index] = mt
		}
	}

	for index, mt := range blUser.Motorcycles {
		if mt.LicensePlate == motorcycleToUpdate.LicensePlate {
			blUser.Motorcycles[index] = motorcycleToUpdate
			break
		}
	}

	updateUser, updateErr := b.UpdateUser(blUser)
	if updateErr != nil {
		return models.User{}, updateErr
	}

	return updateUser, nil
}
