package impl

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"go-forum-api/app/models"
	"go-forum-api/app/repositories"
)

type PostRepository struct {
	db *pgxpool.Pool
}

func CreatePostRepository(db *pgxpool.Pool) repositories.IPostRepository {
	return &PostRepository{db: db}
}

func (repo *PostRepository) Get(id int) (post *models.Post, err error) {
	return
}

func (repo *PostRepository) Update(id int, post *models.Post) (updatedPost *models.Post, err error) {
	return
}
