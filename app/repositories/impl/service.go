package impl

import (
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
	return
}

func (repo *ServiceRepository) Status() (status *models.Status, err error) {
	return
}
