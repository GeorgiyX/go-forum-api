package impl

import (
	"go-forum-api/app/models"
	"go-forum-api/app/repositories"
	"go-forum-api/app/usecases"
)

type UserUseCase struct {
	userRepository repositories.IUserRepository
}

func CreateUserRepository(userRepository repositories.IUserRepository) usecases.IUserUseCase {
	return &UserUseCase{userRepository: userRepository}
}

func (usecase *UserUseCase) Get(nickname *string) (user *models.User, err error) {
	return
}

func (usecase *UserUseCase) All() (users *[]models.User, err error) {
	return
}

func (usecase *UserUseCase) Create(user *models.User) (err error) {
	return
}

func (usecase *UserUseCase) Update(user *models.User) (err error) {
	return
}
