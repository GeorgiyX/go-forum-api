package usecases

import (
	"go-forum-api/app/models"
)

type IServiceUseCase interface {
	Clear() (err error)
	Status() (status *models.Status, err error)
}
