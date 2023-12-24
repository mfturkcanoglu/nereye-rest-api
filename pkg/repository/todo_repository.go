package repository

import (
	"log"

	"github.com/mfturkcan/nereye-rest-api/pkg/model"
	"gorm.io/gorm"
)

type TodoRepository interface{
	GetAllTodos() []model.Todo
}

type todoRepository struct {
	logger *log.Logger
	db *gorm.DB
}

func NewTodoRepository(logger *log.Logger, db *gorm.DB) *todoRepository {
	repo := &todoRepository{
		logger: logger,
		db: db,
	}
	return repo
}