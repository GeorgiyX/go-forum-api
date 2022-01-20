package repositories

import (
	"go-forum-api/app/models"
)

type IThreadRepository interface {
	GetBySlug(slug string) (thread *models.Thread, err error)
	GetByID(id int) (thread *models.Thread, err error)
	UpdateBySlug(thread *models.Thread) (updatedThread *models.Thread, err error)
	UpdateByID(thread *models.Thread) (updatedThread *models.Thread, err error)
}
