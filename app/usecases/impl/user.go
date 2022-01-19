package impl

import (
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"go-forum-api/app/models"
	"go-forum-api/app/repositories"
	"go-forum-api/app/usecases"
	"go-forum-api/utils/errors"
)

type UserUseCase struct {
	userRepository repositories.IUserRepository
}

func CreateUserUseCase(userRepository repositories.IUserRepository) usecases.IUserUseCase {
	return &UserUseCase{userRepository: userRepository}
}

func (usecase *UserUseCase) Get(nickname *string) (user *models.User, err error) {
	user, err = usecase.userRepository.Get(nickname)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = errors.ErrUserNotFound
		} else {
			err = errors.ErrInternalServer
		}
	}
	return
}

func (usecase *UserUseCase) All() (users *[]models.User, err error) {
	return
}

func (usecase *UserUseCase) Create(user *models.User) (err error) {
	err = usecase.userRepository.Create(user)

	if err != nil {
		pgconErr, ok := err.(*pgconn.PgError)
		if ok && pgconErr.SQLState() == errors.SQL23505 {
			err = errors.ErrUserCreate
		} else {
			err = errors.ErrInternalServer
		}
	}
	return
}

func (usecase *UserUseCase) Update(user *models.User) (err error) {
	err = usecase.userRepository.Update(user)

	if err != nil {
		pgconErr, ok := err.(*pgconn.PgError)
		if ok && pgconErr.SQLState() == errors.SQL23505 {
			err = errors.ErrUserUpdate
		} else {
			err = errors.ErrInternalServer
		}
	}
	return
}
