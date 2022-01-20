package impl

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"go-forum-api/app/models"
	"go-forum-api/app/repositories"
	"time"
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
	query := "SELECT id, slug, title, \"user\", posts, threads FROM forums WHERE slug = $1"
	row := repo.db.QueryRow(context.Background(), query, slug)

	forum = &models.Forum{}
	err = row.Scan(&forum.ID, &forum.Slug, &forum.Title, &forum.User, &forum.Posts, &forum.Threads)
	return
}

func (repo *ForumRepository) GetUsers(slug string, params *models.ForumQueryParams) (users []*models.User, err error) {
	return
}

func (repo *ForumRepository) CreateThread(thread *models.Thread) (createdThread *models.Thread, err error) {
	query := "INSERT INTO threads (slug, author, forum, title, message, created) VALUES " +
		"($1, (SELECT nickname FROM users WHERE nickname = $2), " +
		"(SELECT slug FROM forums WHERE slug = $3), $4, $5, $6) " +
		"RETURNING id, COALESCE(slug, ''), author, forum, title, message, created, votes"

	var row pgx.Row
	if thread.Slug == "" {
		row = repo.db.QueryRow(context.Background(), query, nil, thread.Author, thread.Forum,
			thread.Title, thread.Message, thread.Created)

	} else {
		row = repo.db.QueryRow(context.Background(), query, thread.Slug, thread.Author, thread.Forum,
			thread.Title, thread.Message, thread.Created)
	}

	createdThread = &models.Thread{}
	err = row.Scan(&createdThread.ID, &createdThread.Slug, &createdThread.Author, &createdThread.Forum,
		&createdThread.Title, &createdThread.Message, &createdThread.Created, &createdThread.Votes)
	return
}

func (repo *ForumRepository) GetThreads(slug string, params *models.ForumQueryParams) (threads []*models.Thread, err error) {
	// TODO возвможно конкатенация ест перфоманс
	query := "SELECT id, slug, author, forum, title, message, created, votes FROM threads WHERE forum = $1"
	if !params.Since.Equal(time.Time{}) {
		if params.Desc {
			query += " AND created <= $2 ORDER BY created DESC LIMIT $3"
		} else {
			query += " AND created >= $2 ORDER BY created LIMIT $3"
		}
	} else {
		if params.Desc {
			query += " ORDER BY created DESC LIMIT $2"
		} else {
			query += " ORDER BY created LIMIT $2"
		}
	}

	var rows pgx.Rows
	if !params.Since.Equal(time.Time{}) {
		rows, err = repo.db.Query(context.Background(), query, slug, params.Since, params.Limit)
	} else {
		rows, err = repo.db.Query(context.Background(), query, slug, params.Limit)
	}

	defer rows.Close()
	if err != nil {
		return
	}

	threads = make([]*models.Thread, 0)
	for rows.Next() {
		thread := &models.Thread{}
		err = rows.Scan(&thread.ID, &thread.Slug, &thread.Author, &thread.Forum,
			&thread.Title, &thread.Message, &thread.Created, &thread.Votes)
		if err != nil {
			threads = nil
			return
		}
		threads = append(threads, thread)
	}

	return
}
