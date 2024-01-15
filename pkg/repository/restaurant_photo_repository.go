package repository

import (
	"database/sql"
	"log"

	"github.com/mfturkcan/nereye-rest-api/pkg/model"
	"github.com/mfturkcan/nereye-rest-api/pkg/query"
)

type RestaurantPhotoRepository interface {
	CreatePhoto(photoCreate *model.RestaurantPhotoCreate) error
	GetAll() ([]*model.RestaurantPhotoGet, error)
}

type CustomRestaurantPhotoRepository struct {
	logger *log.Logger
	db     *sql.DB
}

func NewRestaurantPhotoRepository(logger *log.Logger, db *sql.DB) *CustomRestaurantPhotoRepository {
	repo := &CustomRestaurantPhotoRepository{
		logger: logger,
		db:     db,
	}
	return repo
}

func (repo *CustomRestaurantPhotoRepository) GetAll(restaurantId string) ([]*model.RestaurantPhotoGet, error) {
	var rows *sql.Rows
	var err error
	if restaurantId == "" {
		rows, err = repo.db.Query(query.RestaurantPhotoSelectQuery())

	} else {
		rows, err = repo.db.Query(query.RestaurantPhotoSelectQueryByRestaurantId(restaurantId))
	}

	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	defer rows.Close()

	photos := []*model.RestaurantPhotoGet{}

	for rows.Next() {
		photo := &model.RestaurantPhotoGet{}
		err := rows.Scan(
			&photo.Name,
			&photo.Url,
			&photo.ParentID,
			&photo.PhotoOrder,
		)

		if err != nil {
			repo.logger.Println(err)
			return nil, err
		}
		photos = append(photos, photo)
	}
	return photos, nil
}

/*
func (repo *CustomRestaurantPhotoRepository) CreatePhoto(customerCreate *model.RestaurantPhotoCreate) error {
	username := customerCreate.CreateRandomCustomerUsername()
	var userId string
	err := repo.db.QueryRow(
		query.UserInsertQueryWithReturn(),
		username, customerCreate.Email, customerCreate.PhoneNumber, customerCreate.FullName, customerCreate.Surname,
	).Scan(&userId)

	if err != nil {
		repo.logger.Println(
			"Error occured while creating customer' user",
			err,
		)
		return err
	}

	_, err = repo.db.Exec(
		query.CustomerInsertQuery(),
		customerCreate.CompanyName, customerCreate.CustomerType, userId,
	)

	if err != nil {
		repo.logger.Println(
			"Error occured while creating customer' customer",
			err,
		)
		return err
	}

	return nil
}
*/
