package repository

import (
	"log"

	"gorm.io/gorm"
)

type Repository struct {
	logger *log.Logger
	db *gorm.DB
}

func NewRepository(logger *log.Logger, db *gorm.DB) *Repository {
	repo := &Repository{
		logger: logger,
		db: db,
	}
	return repo
}