package repository

import (
	"database/sql"
	"log"

	"github.com/mfturkcan/nereye-rest-api/pkg/model"
	"github.com/mfturkcan/nereye-rest-api/pkg/query"
)

type CustomCategoryRepository struct {
	logger *log.Logger
	db     *sql.DB
}

func NewCategoryRepository(logger *log.Logger, db *sql.DB) *CustomCategoryRepository {
	repo := &CustomCategoryRepository{
		logger: logger,
		db:     db,
	}
	return repo
}

func (repo *CustomCategoryRepository) GetAll() ([]*model.CategoryGet, error) {
	rows, err := repo.db.Query(query.CategoryGetQuery())

	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	defer rows.Close()

	categories := []*model.CategoryGet{}

	for rows.Next() {
		category := &model.CategoryGet{}
		err := rows.Scan(
			&category.Id,
			&category.Category,
			&category.PhotoUrl,
		)

		if err != nil {
			repo.logger.Println(err)
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (repo *CustomCategoryRepository) Create(create *model.CategoryCreate) error {
	_, err := repo.db.Exec(
		query.Category_InsertQuery(),
		create.Category,
		create.PhotoUrl,
	)

	if err != nil {
		repo.logger.Println(
			"Error occured while creating category' category",
			err,
		)
		return err
	}

	return nil
}
