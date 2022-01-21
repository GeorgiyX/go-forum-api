package repositories

import (
	"go-forum-api/app/models"
)

type IPostRepository interface {
	Get(id int) (post *models.Post, err error)
	Update(post *models.Post) (updatedPost *models.Post, err error)
}
