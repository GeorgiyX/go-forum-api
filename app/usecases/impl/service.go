package impl

import (
	"go-forum-api/app/models"
	"go-forum-api/app/repositories"
	"go-forum-api/app/usecases"
)

type ServiceUseCase struct {
	serviceRepository repositories.IServiceRepository
}

func CreateServiceUseCase(serviceRepository repositories.IServiceRepository) usecases.IServiceUseCase {
	return &ServiceUseCase{serviceRepository: serviceRepository}
}

func (usecase *ServiceUseCase) Clear() (err error) {
	return
}

func (usecase *ServiceUseCase) Status() (status *models.Status, err error) {
	return
}
