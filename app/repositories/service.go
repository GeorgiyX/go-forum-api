package repositories

import (
	"go-forum-api/app/models"
)

type IServiceRepository interface {
	Clear() (err error)
	Status() (status *models.Status, err error)
}
