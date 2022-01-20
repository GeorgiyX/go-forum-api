package usecases

import "go-forum-api/app/models"

type IForumUseCase interface {
	Create(forum *models.Forum) (createdForum *models.Forum, err error)
	Get(slug string) (forum *models.Forum, err error)
	GetUsers(slug string, params *models.ForumQueryParams) (users []*models.User, err error)
	CreateThread(thread *models.Thread) (createdThread *models.Thread, err error)
	GetThreads(slug string, params *models.ForumQueryParams) (threads []*models.Thread, err error)
}
