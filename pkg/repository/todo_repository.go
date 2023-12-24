package repository

import (
	"database/sql"
	"log"

	"github.com/mfturkcan/nereye-rest-api/pkg/model"
)

type TodoRepository interface{
	GetAllTodos() []model.Todo
}

type todoRepository struct {
	logger *log.Logger
	db *sql.DB
}

func NewTodoRepository(logger *log.Logger, db *sql.DB) *todoRepository {
	repo := &todoRepository{
		logger: logger,
		db: db,
	}
	return repo
}