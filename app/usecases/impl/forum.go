package impl

import (
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"go-forum-api/app/models"
	"go-forum-api/app/repositories"
	"go-forum-api/app/usecases"
	"go-forum-api/utils/errors"
)

type ForumUseCase struct {
	forumRepository repositories.IForumRepository
}

func CreateForumUseCase(forumRepository repositories.IForumRepository) usecases.IForumUseCase {
	return &ForumUseCase{forumRepository: forumRepository}
}

func (usecase *ForumUseCase) Create(forum *models.Forum) (createdForum *models.Forum, err error) {
	createdForum, err = usecase.forumRepository.Create(forum)

	if err != nil {
		pgconErr, ok := err.(*pgconn.PgError)
		if ok {
			switch pgconErr.SQLState() {
			case errors.SQL23502:
				err = errors.ErrForumUserNotFound
				createdForum = nil
				return

			case errors.SQL23505:
				createdForum, err = usecase.forumRepository.Get(forum.Slug)
				if err != nil {
					err = errors.ErrInternalServer
					return
				}
				err = errors.ErrForumAlreadyExists
				return

			default:
				err = errors.ErrInternalServer
			}
		}
	}

	return
}

func (usecase *ForumUseCase) Get(slug string) (forum *models.Forum, err error) {
	forum, err = usecase.forumRepository.Get(slug)

	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.ErrForumNotFound
		} else {
			err = errors.ErrInternalServer
		}
	}

	return
}

func (usecase *ForumUseCase) GetUsers(slug string, params *models.ForumGetUsersQueryParams) (users []*models.User, err error) {
	return
}
