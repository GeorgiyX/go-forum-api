package impl

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"go-forum-api/app/models"
	"go-forum-api/app/repositories"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func CreateUserRepository(db *pgxpool.Pool) repositories.IUserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) Get(nickname *string) (user *models.User, err error) {
	user = &models.User{}
	query := "SELECT nickname, fullname, about, email FROM users WHERE nickname = $1"
	row := repo.db.QueryRow(context.Background(), query, *nickname)
	err = row.Scan(&user.NickName, &user.FullName, &user.About, &user.Email)
	return
}

func (repo *UserRepository) All() (users *[]models.User, err error) {
	return
}

func (repo *UserRepository) Create(user *models.User) (err error) {
	query := "INSERT INTO users (nickname, fullname, about, email) VALUES ($1, $2, $3, $4)"
	_, err = repo.db.Exec(context.Background(), query, user.NickName, user.FullName, user.About, user.Email)
	return
}

func (repo *UserRepository) Update(user *models.User) (updatedUser *models.User, err error) {
	query := "UPDATE users SET fullname = $1, about = $2, email = $3 WHERE nickname = $4 " +
		"RETURNING nickname, fullname, about, email"

	row := repo.db.QueryRow(context.Background(), query, user.FullName, user.About, user.Email, user.NickName)
	updatedUser = &models.User{}
	err = row.Scan(&updatedUser.NickName, &updatedUser.FullName, &updatedUser.About, &updatedUser.Email)
	return
}

func (repo *UserRepository) GetUsersByUserNicknameOrEmail(user *models.User) (users []*models.User, err error) {
	query := "SELECT nickname, fullname, about, email FROM users WHERE nickname = $1 OR email = $2"
	rows, err := repo.db.Query(context.Background(), query, user.NickName, user.Email)
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		conflictUser := &models.User{}
		err = rows.Scan(&conflictUser.NickName, &conflictUser.FullName, &conflictUser.About, &conflictUser.Email)
		if err != nil {
			users = nil
			return
		}
		users = append(users, conflictUser)
	}

	return
}
