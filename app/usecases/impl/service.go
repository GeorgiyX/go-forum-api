package impl

import (
	"go-forum-api/app/models"
	"go-forum-api/app/repositories"
	"go-forum-api/app/usecases"
	"go-forum-api/utils/errors"
)

type ServiceUseCase struct {
	serviceRepository repositories.IServiceRepository
}

func CreateServiceUseCase(serviceRepository repositories.IServiceRepository) usecases.IServiceUseCase {
	return &ServiceUseCase{serviceRepository: serviceRepository}
}

func (usecase *ServiceUseCase) Clear() (err error) {
	err = usecase.serviceRepository.Clear()
	if err != nil {
		err = errors.ErrInternalServer
		return
	}
	return
}

func (usecase *ServiceUseCase) Status() (status *models.Status, err error) {
	status, err = usecase.serviceRepository.Status()
	if err != nil {
		err = errors.ErrInternalServer
		return
	}
	return
}
