package impl

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"go-forum-api/app/models"
	"go-forum-api/app/repositories"
)

type ThreadRepository struct {
	db *pgxpool.Pool
}

func CreateThreadRepository(db *pgxpool.Pool) repositories.IThreadRepository {
	return &ThreadRepository{db: db}
}

func (repo *ThreadRepository) GetBySlug(slug string) (thread *models.Thread, err error) {
	return
}
func (repo *ThreadRepository) GetByID(id int) (thread *models.Thread, err error) {
	return
}
func (repo *ThreadRepository) UpdateBySlug(thread *models.Thread) (updatedThread *models.Thread, err error) {
	return
}
func (repo *ThreadRepository) UpdateByID(thread *models.Thread) (updatedThread *models.Thread, err error) {
	return
}
