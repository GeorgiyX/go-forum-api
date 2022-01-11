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

func (repo *UserRepository) Update(user *models.User) (err error) {
	return
}
