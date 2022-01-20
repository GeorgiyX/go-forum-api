package repositories

import "go-forum-api/app/models"

type IUserRepository interface {
	Get(nickname *string) (user *models.User, err error)
	All() (users *[]models.User, err error)
	Create(user *models.User) (err error)
	Update(user *models.User) (updatedUser *models.User, err error)
	GetUsersByUserNicknameOrEmail(user *models.User) (users []*models.User, err error)
}
