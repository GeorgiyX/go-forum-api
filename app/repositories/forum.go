package repositories

import "go-forum-api/app/models"

type IForumRepository interface {
	Create(forum *models.Forum) (createdForum *models.Forum, err error)
	Get(slug string) (forum *models.Forum, err error)
	GetUsers(slug string, params *models.ForumGetUsersQueryParams) (users []*models.User, err error)
}
