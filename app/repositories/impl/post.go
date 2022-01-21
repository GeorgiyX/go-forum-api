package impl

import (
	"context"
	"github.com/jackc/pgx/v4"
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
	query := "SELECT id, COALESCE(parent, 0), author, forum, thread, created, isEdited, message FROM posts WHERE id = $1"
	row := repo.db.QueryRow(context.Background(), query, id)
	post = &models.Post{}
	err = row.Scan(&post.ID, &post.Parent, &post.Author, &post.Forum, &post.Thread,
		&post.Created, &post.IsEdited, &post.Message)
	return
}

func (repo *PostRepository) Update(post *models.Post) (updatedPost *models.Post, err error) {
	query := "UPDATE posts SET message = COALESCE(NULLIF($1, ''), message), " +
		"isEdited = CASE WHEN (isEdited = TRUE OR (isEdited = FALSE AND $1 IS NOT NULL AND $1 <> message)) " +
		"THEN TRUE ELSE FALSE END WHERE id = $2 " +
		"RETURNING id, COALESCE(parent, 0), author, forum, thread, created, isEdited, message"
	var row pgx.Row
	if post.Message != "" {
		row = repo.db.QueryRow(context.Background(), query, post.Message, post.ID)
	} else {
		row = repo.db.QueryRow(context.Background(), query, nil, post.ID)
	}

	updatedPost = &models.Post{}
	err = row.Scan(&updatedPost.ID, &updatedPost.Parent, &updatedPost.Author, &updatedPost.Forum,
		&updatedPost.Thread, &updatedPost.Created, &updatedPost.IsEdited, &updatedPost.Message)
	return
}
