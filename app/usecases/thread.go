package usecases

import (
	"go-forum-api/app/models"
)

type IThreadUseCase interface {
	Get(slugOrId string) (thread *models.Thread, err error)
	Update(slugOrId string, thread *models.Thread) (updatedThread *models.Thread, err error)
}
