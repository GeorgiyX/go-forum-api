package impl

import (
	"go-forum-api/app/models"
	"go-forum-api/app/repositories"
	"go-forum-api/app/usecases"
)

type ForumUseCase struct {
	forumRepository repositories.IForumRepository
}

func CreateForumUseCase(forumRepository repositories.IForumRepository) usecases.IForumUseCase {
	return &ForumUseCase{forumRepository: forumRepository}
}

func (repo *ForumUseCase) Create(forum *models.Forum) (createdForum *models.Forum, err error) {
	return
}

func (repo *ForumUseCase) Get(slug string) (forum *models.Forum, err error) {
	return
}

func (repo *ForumUseCase) GetUsers(slug string, params *models.ForumGetUsersQueryParams) (users []*models.User, err error) {
	return
}
