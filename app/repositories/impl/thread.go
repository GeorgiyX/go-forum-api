package impl

import (
	"context"
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
	query := "SELECT id, slug, author, forum, title, message, created, votes FROM threads WHERE slug = $1"
	row := repo.db.QueryRow(context.Background(), query, slug)

	thread = &models.Thread{}
	err = row.Scan(&thread.ID, &thread.Slug, &thread.Author, &thread.Forum,
		&thread.Title, &thread.Message, &thread.Created, &thread.Votes)
	return
}
func (repo *ThreadRepository) GetByID(id int) (thread *models.Thread, err error) {
	query := "SELECT id, slug, author, forum, title, message, created, votes FROM threads WHERE id = $1"
	row := repo.db.QueryRow(context.Background(), query, id)

	thread = &models.Thread{}
	err = row.Scan(&thread.ID, &thread.Slug, &thread.Author, &thread.Forum,
		&thread.Title, &thread.Message, &thread.Created, &thread.Votes)
	return
}
func (repo *ThreadRepository) UpdateBySlug(thread *models.Thread) (updatedThread *models.Thread, err error) {
	query := "UPDATE threads SET title = COALESCE(NULLIF($1, ''), title), " +
		"message = COALESCE(NULLIF($2, ''), message) WHERE slug = $3 " +
		"RETURNING id, slug, author, forum, title, message, created, votes"

	row := repo.db.QueryRow(context.Background(), query, thread.Title, thread.Message, thread.Slug)
	updatedThread = &models.Thread{}
	err = row.Scan(&updatedThread.ID, &updatedThread.Slug, &updatedThread.Author, &updatedThread.Forum,
		&updatedThread.Title, &updatedThread.Message, &updatedThread.Created, &updatedThread.Votes)
	return
}
func (repo *ThreadRepository) UpdateByID(thread *models.Thread) (updatedThread *models.Thread, err error) {
	query := "UPDATE threads SET title = COALESCE(NULLIF($1, ''), title), " +
		"message = COALESCE(NULLIF($2, ''), message) WHERE id = $3 " +
		"RETURNING id, slug, author, forum, title, message, created, votes"

	row := repo.db.QueryRow(context.Background(), query, thread.Title, thread.Message, thread.ID)
	updatedThread = &models.Thread{}
	err = row.Scan(&updatedThread.ID, &updatedThread.Slug, &updatedThread.Author, &updatedThread.Forum,
		&updatedThread.Title, &updatedThread.Message, &updatedThread.Created, &updatedThread.Votes)
	return
}
