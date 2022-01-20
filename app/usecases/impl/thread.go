package impl

import (
	"go-forum-api/app/models"
	"go-forum-api/app/repositories"
	"go-forum-api/app/usecases"
)

type ThreadUseCase struct {
	threadRepository repositories.IThreadRepository
}

func CreateThreadUseCase(threadRepository repositories.IThreadRepository) usecases.IThreadUseCase {
	return &ThreadUseCase{threadRepository: threadRepository}
}

func (usecase *ThreadUseCase) GetBySlug(slug string) (thread *models.Thread, err error) {
	return
}
func (usecase *ThreadUseCase) GetByID(id int) (thread *models.Thread, err error) {
	return
}
func (usecase *ThreadUseCase) UpdateBySlug(thread *models.Thread) (updatedThread *models.Thread, err error) {
	return
}
func (usecase *ThreadUseCase) UpdateByID(thread *models.Thread) (updatedThread *models.Thread, err error) {
	return
}
