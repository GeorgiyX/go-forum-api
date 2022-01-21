package usecases

import (
	"go-forum-api/app/models"
)

type IThreadUseCase interface {
	Get(slugOrId string) (thread *models.Thread, err error)
	Update(slugOrId string, thread *models.Thread) (updatedThread *models.Thread, err error)
	Vote(slugOrId string, vote *models.Vote) (thread *models.Thread, err error)
	CreatePosts(slugOrId string, posts []*models.Post) (createdPosts []*models.Post, err error)
}
