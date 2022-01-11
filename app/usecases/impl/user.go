package impl

import (
	"fmt"
	"go-forum-api/app/models"
	"go-forum-api/app/repositories"
	"go-forum-api/app/usecases"
)

type UserUseCase struct {
	userRepository repositories.IUserRepository
}

func CreateUserUseCase(userRepository repositories.IUserRepository) usecases.IUserUseCase {
	return &UserUseCase{userRepository: userRepository}
}

func (usecase *UserUseCase) Get(nickname *string) (user *models.User, err error) {
	user, err = usecase.userRepository.Get(nickname)
	// todo выяснить как получить детали об ошикбке для 409
	if err != nil {
		switch err {
		case :
			
		
		}
		fmt.Printf("Ошибка при get user: %T [%+v]", err, err)
	}
	return
}

func (usecase *UserUseCase) All() (users *[]models.User, err error) {
	return
}

func (usecase *UserUseCase) Create(user *models.User) (err error) {
	err = usecase.userRepository.Create(user)
	// todo выяснить как получить детали об ошикбке для 409
	if err != nil {
		fmt.Printf("Ошибка при create user: %T [%+v]", err, err)
		return err
	}
	return
}

func (usecase *UserUseCase) Update(user *models.User) (err error) {
	return
}
