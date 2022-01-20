package impl

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"go-forum-api/app/models"
	"go-forum-api/app/repositories"
)

type ForumRepository struct {
	db *pgxpool.Pool
}

func CreateForumRepository(db *pgxpool.Pool) repositories.IForumRepository {
	return &ForumRepository{db: db}
}

func (repo *ForumRepository) Create(forum *models.Forum) (createdForum *models.Forum, err error) {
	// Для "не точного" равенства nickname подзапрос
	query := "INSERT INTO forums (slug, title, \"user\") VALUES ($1, $2, " +
		"(SELECT nickname FROM users WHERE nickname = $3)) RETURNING slug, title, \"user\", posts, threads"
	row := repo.db.QueryRow(context.Background(), query, forum.Slug, forum.Title, forum.User)
	createdForum = &models.Forum{}
	err = row.Scan(&createdForum.Slug, &createdForum.Title, &createdForum.User,
		&createdForum.Posts, &createdForum.Threads)
	return
}

func (repo *ForumRepository) Get(slug string) (forum *models.Forum, err error) {
	query := "SELECT slug, title, \"user\", posts, threads FROM forums WHERE slug = $1"
	row := repo.db.QueryRow(context.Background(), query, slug)
	forum = &models.Forum{}
	err = row.Scan(&forum.Slug, &forum.Title, &forum.User, &forum.Posts, &forum.Threads)
	return
}

func (repo *ForumRepository) GetUsers(slug string, params *models.ForumGetUsersQueryParams) (users []*models.User, err error) {
	return
}

func (repo *ForumRepository) CreateThread(thread *models.Thread) (createdThread *models.Thread, err error) {
	return
}

func (repo *ForumRepository) GetThreads(slug string) (threads []*models.Thread, err error) {
	return
}
