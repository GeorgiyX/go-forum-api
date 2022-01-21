package impl

import (
	"go-forum-api/app/models"
	"go-forum-api/app/repositories"
	"go-forum-api/app/usecases"
)

type PostUseCase struct {
	postRepository repositories.IPostRepository
}

func CreatePostUseCase(postRepository repositories.IPostRepository) usecases.IPostUseCase {
	return &PostUseCase{postRepository: postRepository}
}

func (usecase *PostUseCase) Get(id int) (post *models.Post, err error) {
	return
}

func (usecase *PostUseCase) Update(id int, post *models.Post) (updatedPost *models.Post, err error) {
	return
}

