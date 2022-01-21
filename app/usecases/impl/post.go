package impl

import (
	"github.com/jackc/pgx/v4"
	"go-forum-api/app/models"
	"go-forum-api/app/repositories"
	"go-forum-api/app/usecases"
	"go-forum-api/utils/constants"
	"go-forum-api/utils/errors"
)

type PostUseCase struct {
	postRepository repositories.IPostRepository
	forumUseCase   usecases.IForumUseCase
	userUseCase    usecases.IUserUseCase
	threadUseCase  usecases.IThreadUseCase
}

func CreatePostUseCase(postRepository repositories.IPostRepository,
	forumUseCase usecases.IForumUseCase,
	userUseCase usecases.IUserUseCase,
	threadUseCase usecases.IThreadUseCase) usecases.IPostUseCase {
	return &PostUseCase{
		postRepository: postRepository,
		forumUseCase:   forumUseCase,
		userUseCase:    userUseCase,
		threadUseCase:  threadUseCase,
	}
}

func (usecase *PostUseCase) Get(id int, details []string) (postDetailed *models.PostDetailed, err error) {
	postDetailed = &models.PostDetailed{}
	postDetailed.Post, err = usecase.postRepository.Get(id)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.ErrPostNotFound
			return
		}
		err = errors.ErrInternalServer
		return
	}

	for _, detailType := range details {
		switch detailType {
		case constants.PostUser:
			postDetailed.Author, err = usecase.userUseCase.Get(&postDetailed.Post.Author)
			if err != nil {
				postDetailed = nil
				return
			}
		case constants.PostThread:
			postDetailed.Thread, err = usecase.threadUseCase.Get(string(rune(postDetailed.Post.Thread)))
			if err != nil {
				postDetailed = nil
				return
			}
		case constants.PostForum:
			postDetailed.Forum, err = usecase.forumUseCase.Get(postDetailed.Post.Forum)
			if err != nil {
				postDetailed = nil
				return
			}
		default:
			postDetailed = nil
			err = errors.ErrBadRequest.SetDetails("неверные query параметры")
			return
		}
	}

	return
}

func (usecase *PostUseCase) Update(post *models.Post) (updatedPost *models.Post, err error) {
	updatedPost, err = usecase.postRepository.Update(post)

	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.ErrPostNotFound
			return
		}
		err = errors.ErrInternalServer
		return
	}

	return
}
