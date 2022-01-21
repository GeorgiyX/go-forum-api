package usecases

import (
	"go-forum-api/app/models"
)

type IPostUseCase interface {
	Get(id int, details []string) (postDetailed *models.PostDetailed, err error)
	Update(post *models.Post) (updatedPost *models.Post, err error)
}
