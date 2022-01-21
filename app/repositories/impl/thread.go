package impl

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"go-forum-api/app/models"
	"go-forum-api/app/repositories"
	"go-forum-api/utils/constants"
	"time"
)

type ThreadRepository struct {
	db *pgxpool.Pool
}

func CreateThreadRepository(db *pgxpool.Pool) repositories.IThreadRepository {
	return &ThreadRepository{db: db}
}

func (repo *ThreadRepository) GetBySlug(slug string) (thread *models.Thread, err error) {
	query := "SELECT id, COALESCE(slug, ''), author, forum, title, message, created, votes FROM threads WHERE slug = $1"
	row := repo.db.QueryRow(context.Background(), query, slug)

	thread = &models.Thread{}
	err = row.Scan(&thread.ID, &thread.Slug, &thread.Author, &thread.Forum,
		&thread.Title, &thread.Message, &thread.Created, &thread.Votes)
	return
}
func (repo *ThreadRepository) GetByID(id int) (thread *models.Thread, err error) {
	query := "SELECT id, COALESCE(slug, ''), author, forum, title, message, created, votes FROM threads WHERE id = $1"
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
		"RETURNING id, COALESCE(slug, ''), author, forum, title, message, created, votes"

	row := repo.db.QueryRow(context.Background(), query, thread.Title, thread.Message, thread.ID)
	updatedThread = &models.Thread{}
	err = row.Scan(&updatedThread.ID, &updatedThread.Slug, &updatedThread.Author, &updatedThread.Forum,
		&updatedThread.Title, &updatedThread.Message, &updatedThread.Created, &updatedThread.Votes)
	return
}

func (repo *ThreadRepository) VoteBySlug(slug string, vote *models.Vote) (err error) {
	query := "INSERT INTO votes (nickname, thread, value) VALUES ($1, (SELECT id FROM threads WHERE slug=$2), $3) " +
		"ON CONFLICT (nickname, thread) DO UPDATE SET value = $3"
	_, err = repo.db.Exec(context.Background(), query, vote.NickName, slug, vote.Voice)
	return
}

func (repo *ThreadRepository) VoteByID(id int, vote *models.Vote) (err error) {
	query := "INSERT INTO votes (nickname, thread, value) VALUES ($1, $2, $3) " +
		"ON CONFLICT (nickname, thread) DO UPDATE SET value = $3"
	_, err = repo.db.Exec(context.Background(), query, vote.NickName, id, vote.Voice)
	return
}

func (repo *ThreadRepository) CreatePosts(threadId int, forumSlug string, posts []*models.Post) (createdPosts []*models.Post, err error) {
	query := "INSERT INTO posts(parent, author, forum, thread, message, created) " +
		"VALUES ($1, $2, $3, $4, $5, $6) " +
		"RETURNING id, COALESCE(parent, 0), author, forum, thread, created, isEdited, message"

	ctx := context.Background()
	tx, err := repo.db.Begin(ctx)
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			trErr := tx.Rollback(ctx)
			if trErr != nil {
				err = trErr
			}
		} else {
			trErr := tx.Commit(ctx)
			if trErr != nil {
				err = trErr
			}
		}
	}()

	batch := new(pgx.Batch)
	createdTime := time.Now()

	for _, post := range posts {
		if post.Parent == 0 {
			batch.Queue(query, nil, post.Author, forumSlug, threadId, post.Message, createdTime)
		} else {
			batch.Queue(query, post.Parent, post.Author, forumSlug, threadId, post.Message, createdTime)
		}
	}

	batchRes := tx.SendBatch(ctx, batch)
	defer func() {
		batchErr := batchRes.Close()
		if batchErr != nil {
			err = batchErr
		}
	}()

	createdPosts = make([]*models.Post, 0)

	for i := 0; i < batch.Len(); i++ {
		createdPost := &models.Post{}

		row := batchRes.QueryRow()
		err = row.Scan(&createdPost.ID, &createdPost.Parent, &createdPost.Author, &createdPost.Forum,
			&createdPost.Thread, &createdPost.Created, &createdPost.IsEdited, &createdPost.Message)

		if err != nil {
			createdPosts = nil
			return
		}

		createdPosts = append(createdPosts, createdPost)
	}

	return
}

func (repo *ThreadRepository) GetPosts(threadId int, params *models.PostsQueryParams) (posts []*models.Post, err error) {
	var rows pgx.Rows

	if params.Since == 0 {
		if params.Desc {
			rows, err = repo.db.Query(context.Background(), constants.DescNoSincePostQuery[params.Sort],
				threadId, params.Limit)
		} else {
			rows, err = repo.db.Query(context.Background(), constants.AscNoSincePostQuery[params.Sort],
				threadId, params.Limit)
		}
	} else {
		if params.Desc {
			rows, err = repo.db.Query(context.Background(), constants.DescSincePostQuery[params.Sort],
				threadId, params.Since, params.Limit)
		} else {
			rows, err = repo.db.Query(context.Background(), constants.AscSincePostQuery[params.Sort],
				threadId, params.Since, params.Limit)
		}
	}

	defer rows.Close()
	if err != nil {
		return
	}

	posts = make([]*models.Post, 0)
	for rows.Next() {
		post := &models.Post{}
		err = rows.Scan(&post.ID, &post.Parent, &post.Author, &post.Forum,
			&post.Thread, &post.Created, &post.IsEdited, &post.Message)
		if err != nil {
			posts = nil
			return
		}
		posts = append(posts, post)
	}
	return
}
