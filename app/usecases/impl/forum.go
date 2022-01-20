package impl

import (
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"go-forum-api/app/models"
	"go-forum-api/app/repositories"
	"go-forum-api/app/usecases"
	"go-forum-api/utils/errors"
)

type ForumUseCase struct {
	forumRepository  repositories.IForumRepository
	threadRepository repositories.IThreadRepository
}

func CreateForumUseCase(forumRepository repositories.IForumRepository,
	threadRepository repositories.IThreadRepository) usecases.IForumUseCase {
	return &ForumUseCase{forumRepository: forumRepository, threadRepository: threadRepository}
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

func (usecase *ForumUseCase) GetUsers(slug string, params *models.ForumQueryParams) (users []*models.User, err error) {
	return
}

func (usecase *ForumUseCase) CreateThread(thread *models.Thread) (createdThread *models.Thread, err error) {
	createdThread, err = usecase.forumRepository.CreateThread(thread)

	if err != nil {
		pgconErr, ok := err.(*pgconn.PgError)
		if ok {
			switch pgconErr.SQLState() {
			case errors.SQL23502:
				err = errors.ErrThreadUserOrForumNotFound
				createdThread = nil
				return

			case errors.SQL23505:
				createdThread, err = usecase.threadRepository.GetBySlug(thread.Slug)
				if err != nil {
					err = errors.ErrInternalServer
					return
				}
				err = errors.ErrThreadAlreadyExists
				return

			default:
				err = errors.ErrInternalServer
			}
		}
	}

	return
}

func (usecase *ForumUseCase) GetThreads(slug string, params *models.ForumQueryParams) (threads []*models.Thread, err error) {
	threads, err = usecase.forumRepository.GetThreads(slug, params)
	if err != nil {
		fmt.Printf("error get threads: %v", err)
		err = errors.ErrInternalServer
		return
	}

	if len(threads) == 0 {
		_, err = usecase.Get(slug)
		if err != nil {
			fmt.Printf("error get threads: %v", err)
			return
		}
		//TODO стоит добавть пустой слайс если nil не ок
	}

	return
}
