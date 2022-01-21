package impl

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"go-forum-api/app/models"
	"go-forum-api/app/repositories"
)

type ServiceRepository struct {
	db *pgxpool.Pool
}

func CreateServiceRepository(db *pgxpool.Pool) repositories.IServiceRepository {
	return &ServiceRepository{db: db}
}

func (repo *ServiceRepository) Clear() (err error) {
	query := "TRUNCATE users, forums, threads, votes, posts, forum_users"
	_, err = repo.db.Exec(context.Background(), query)
	return
}

func (repo *ServiceRepository) Status() (status *models.Status, err error) {
	queryUsers := "SELECT COUNT(*) FROM users"
	queryForums := "SELECT COUNT(*) FROM forums"
	queryThreads := "SELECT COUNT(*) FROM threads"
	queryPosts := "SELECT COUNT(*) FROM posts"

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

	status = &models.Status{}
	if err = tx.QueryRow(ctx, queryUsers).Scan(&status.User); err != nil {
		status = nil
		return
	}
	if err = tx.QueryRow(ctx, queryForums).Scan(&status.Forum); err != nil {
		status = nil
		return
	}
	if err = tx.QueryRow(ctx, queryThreads).Scan(&status.Thread); err != nil {
		status = nil
		return
	}
	if err = tx.QueryRow(ctx, queryPosts).Scan(&status.Post); err != nil {
		status = nil
		return
	}

	return
}
