package impl

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"go-forum-api/app/models"
	"go-forum-api/app/repositories"
)

type ForumRepository struct {
	db *pgxpool.Pool
}

func CreateForumRepository(db *pgxpool.Pool) repositories.IForumRepository {
	return &ForumRepository{db: db}
}

func (repo *ForumRepository) Create(forum *models.Forum) (createdForum *models.Forum, err error) {
	return
}

func (repo *ForumRepository) Get(slug string) (forum *models.Forum, err error) {
	return
}

func (repo *ForumRepository) GetUsers(slug string, params *models.ForumGetUsersQueryParams) (users []*models.User, err error) {
	return
}
