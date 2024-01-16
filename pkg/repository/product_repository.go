package repository

import (
	"database/sql"
	"log"

	"github.com/mfturkcan/nereye-rest-api/pkg/model"
	"github.com/mfturkcan/nereye-rest-api/pkg/query"
)

type CustomProductRepository struct {
	logger *log.Logger
	db     *sql.DB
}

func NewProductRepository(logger *log.Logger, db *sql.DB) *CustomProductRepository {
	repo := &CustomProductRepository{
		logger: logger,
		db:     db,
	}
	return repo
}

func (repo *CustomProductRepository) GetAll() ([]*model.ProductGetAll, error) {
	rows, err := repo.db.Query(query.ProductGetQuery())

	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	defer rows.Close()

	products := []*model.ProductGetAll{}

	for rows.Next() {
		product := &model.ProductGetAll{}
		err := rows.Scan(
			&product.ID,
			&product.ProductName,
			&product.PhotoUrl,
			&product.AvailableAtStart,
			&product.AvailableAtEnd,
			&product.RestaurantId,
			&product.CategoryId,
		)

		if err != nil {
			repo.logger.Println(err)
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (repo *CustomProductRepository) Create(create *model.ProductCreate) error {
	_, err := repo.db.Exec(
		query.Product_InsertQuery(),
		create.ProductName,
		create.PhotoUrl,
		create.AvailableAtStart,
		create.AvailableAtEnd,
		create.RestaurantId,
		create.CategoryId,
	)

	if err != nil {
		repo.logger.Println(
			"Error occured while creating product' product",
			err,
		)
		return err
	}

	return nil
}
