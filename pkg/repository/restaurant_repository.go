package repository

import (
	"database/sql"
	"log"

	"github.com/mfturkcan/nereye-rest-api/pkg/model"
	"github.com/mfturkcan/nereye-rest-api/pkg/query"
)

type RestaurantRepository interface {
	CreateRestaurant(createDto *model.RestaurantCreate) error
	GetAll() ([]*model.RestaurantCreate, error)
}

type CustomRestaurantRepository struct {
	logger *log.Logger
	db     *sql.DB
}

func NewRestaurantRepository(logger *log.Logger, db *sql.DB) *CustomRestaurantRepository {
	repo := &CustomRestaurantRepository{
		logger: logger,
		db:     db,
	}
	return repo
}

func (repo *CustomRestaurantRepository) CreateRestaurant(create *model.RestaurantCreate) error {
	var addressId string
	err := repo.db.QueryRow(
		query.CreateAddressQueryReturnId(),
		create.Country,
		create.City,
		create.County,
		create.District,
		create.FullAddress,
		create.Latitude,
		create.Longitude,
	).Scan(&addressId)

	if err != nil {
		repo.logger.Println(
			"Error occured while creating restaurant's address",
			err,
		)
		return err
	}

	_, err = repo.db.Exec(
		query.Restaurant_InsertQuery(),
		create.PhotoUrl,
		create.SignName,
		create.AboutUs,
		create.ExtraInfo,
		create.PhoneNumber,
		create.WorkplacePhoneNumber,
		create.AvailableAtStart,
		create.AvailableAtEnd,
		create.WeekendAvailableAtStart,
		create.WeekendAvailableAtEnd,
		addressId,
		create.CustomerId,
	)

	if err != nil {
		repo.logger.Println(
			"Error occured while creating restaurant' restaurant",
			err,
		)
		return err
	}

	return nil
}

func (repo *CustomRestaurantRepository) GetAll(customerId string) ([]*model.RestaurantGet, error) {
	rows, err := repo.db.Query(query.Restaurant_All(customerId))

	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	defer rows.Close()

	return repo.getRestaurantsFromRows(rows)
}

func (repo *CustomRestaurantRepository) getRestaurantsFromRows(rows *sql.Rows) ([]*model.RestaurantGet, error) {
	restaurants := []*model.RestaurantGet{}

	for rows.Next() {
		restaurant := &model.RestaurantGet{}
		err := rows.Scan(
			&restaurant.PhotoUrl,
			&restaurant.SignName,
			&restaurant.AboutUs,
			&restaurant.ExtraInfo,
			&restaurant.PhoneNumber,
			&restaurant.WorkplacePhoneNumber,
			&restaurant.AvailableAtStart,
			&restaurant.AvailableAtEnd,
			&restaurant.WeekendAvailableAtStart,
			&restaurant.WeekendAvailableAtEnd,
			&restaurant.Country,
			&restaurant.City,
			&restaurant.County,
			&restaurant.District,
			&restaurant.FullAddress,
			&restaurant.Latitude,
			&restaurant.Longitude,
		)

		// TODO: Get it from time_util
		restaurant.IsAvailable = true

		if err != nil {
			repo.logger.Println(err)
			return nil, err
		}
		restaurants = append(restaurants, restaurant)
	}
	return restaurants, nil
}
