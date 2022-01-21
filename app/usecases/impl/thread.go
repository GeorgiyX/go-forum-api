package impl

import (
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"go-forum-api/app/models"
	"go-forum-api/app/repositories"
	"go-forum-api/app/usecases"
	"go-forum-api/utils/errors"
	"go-forum-api/utils/validator"
)

type ThreadUseCase struct {
	threadRepository repositories.IThreadRepository
}

func CreateThreadUseCase(threadRepository repositories.IThreadRepository) usecases.IThreadUseCase {
	return &ThreadUseCase{threadRepository: threadRepository}
}

func (usecase *ThreadUseCase) Get(slugOrId string) (thread *models.Thread, err error) {
	v, _ := validator.GetInstance()
	slug, id, err := v.GetSlugOrIdOrErr(slugOrId)
	if err != nil {
		err = errors.ErrBadRequest.SetDetails(err.Error())
		return
	}

	if slug == "" {
		thread, err = usecase.threadRepository.GetByID(id)
	} else {
		thread, err = usecase.threadRepository.GetBySlug(slug)
	}

	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.ErrThreadNotFound
		} else {
			err = errors.ErrInternalServer
		}
	}

	return
}

func (usecase *ThreadUseCase) Update(slugOrId string, thread *models.Thread) (updatedThread *models.Thread, err error) {
	v, _ := validator.GetInstance()
	slug, id, err := v.GetSlugOrIdOrErr(slugOrId)
	if err != nil {
		err = errors.ErrBadRequest.SetDetails(err.Error())
		return
	}

	if slug == "" {
		thread.ID = id
		updatedThread, err = usecase.threadRepository.UpdateByID(thread)
	} else {
		thread.Slug = slug
		updatedThread, err = usecase.threadRepository.UpdateBySlug(thread)
	}

	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.ErrThreadUpdateNotFound
			return
		}
		err = errors.ErrInternalServer
		return
	}

	return
}

func (usecase *ThreadUseCase) Vote(slugOrId string, vote *models.Vote) (thread *models.Thread, err error) {
	v, _ := validator.GetInstance()
	if !v.ValidateVote(vote) {
		err = errors.ErrBadRequest.SetDetails("не верное значение голоса")
		return
	}

	slug, id, err := v.GetSlugOrIdOrErr(slugOrId)
	if err != nil {
		err = errors.ErrBadRequest.SetDetails(err.Error())
		return
	}

	if slug == "" {
		err = usecase.threadRepository.VoteByID(id, vote)
	} else {
		err = usecase.threadRepository.VoteBySlug(slug, vote)
	}

	if err != nil {
		pgconErr, ok := err.(*pgconn.PgError)
		if ok {
			if pgconErr.SQLState() == errors.SQL23503 {
				err = errors.ErrThreadUserOrThreadNotFound
				return
			} else {
				err = errors.ErrInternalServer
				return
			}
		}
	}

	if slug == "" {
		thread, err = usecase.threadRepository.GetByID(id)
	} else {
		thread, err = usecase.threadRepository.GetBySlug(slug)
	}

	if err != nil {
		err = errors.ErrInternalServer
		return
	}

	return
}

func (usecase *ThreadUseCase) CreatePosts(slugOrId string, posts []*models.Post) (createdPosts []*models.Post, err error) {
	thread, err := usecase.Get(slugOrId)
	if err != nil {
		return
	}

	createdPosts, err = usecase.threadRepository.CreatePosts(thread.ID, thread.Forum, posts)
	if err != nil {
		pgconErr, ok := err.(*pgconn.PgError)
		if ok {
			switch pgconErr.SQLState() {
			case errors.SQL23503:
				err = errors.ErrPostUserNotFound
				return

			case errors.P0001:
				err = errors.ErrPostWrongParent
				return

			default:
				err = errors.ErrInternalServer
			}
		}
	}
	
	return
}
