package usecases

import (
	"go-forum-api/app/models"
)

type IThreadUseCase interface {
	GetBySlug(slug string) (thread *models.Thread, err error)
	GetByID(id int) (thread *models.Thread, err error)
	UpdateBySlug(thread *models.Thread) (updatedThread *models.Thread, err error)
	UpdateByID(thread *models.Thread) (updatedThread *models.Thread, err error)
}
