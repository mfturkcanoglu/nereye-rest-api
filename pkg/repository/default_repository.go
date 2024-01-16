package repository

import (
	"database/sql"
	"log"

	"github.com/mfturkcan/nereye-rest-api/pkg/model"
)

type DefaultRepository[T model.DefaultModel] interface {
	GetAll() ([]*T, error)
}

type CustomDefaultRepository struct {
	logger *log.Logger
	db     *sql.DB
}

func NewDefaultRepository(logger *log.Logger, db *sql.DB) *CustomDefaultRepository {
	repo := &CustomDefaultRepository{
		logger: logger,
		db:     db,
	}
	return repo
}
