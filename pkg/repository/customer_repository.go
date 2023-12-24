package repository

import (
	"database/sql"
	"log"

	"github.com/mfturkcan/nereye-rest-api/pkg/model"
)

type CustomerRepository interface {
	CreateCustomer(customerCreate *model.CustomerCreate) error
	GetAll() ([]*model.CustomerCreate, error)
}

type CustomCustomerRepository struct {
	logger *log.Logger
	db     *sql.DB
}

func NewCustomerRepository(logger *log.Logger, db *sql.DB) *CustomCustomerRepository {
	repo := &CustomCustomerRepository{
		logger: logger,
		db:     db,
	}
	return repo
}

func (repo *CustomCustomerRepository) GetAll() ([]*model.UserGet, error) {
	rows, err := repo.db.Query(`
		SELECT
		u.username,
		u.phone_number,
		u.email,
		u.full_name,
		u.surname
		from users u
		order by u.updated_at desc
	`)

	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	defer rows.Close()

	users := []*model.UserGet{}

	for rows.Next() {
		user := &model.UserGet{}
		err := rows.Scan(
			&user.Username,
			&user.PhoneNumber,
			&user.Email,
			&user.FullName,
			&user.Surname)

		if err != nil {
			repo.logger.Println(err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (repo *CustomCustomerRepository) CreateCustomer(customerCreate *model.CustomerCreate) error {
	username := customerCreate.CreateRandomCustomerUsername()
	var userId string
	err := repo.db.QueryRow(
		"INSERT INTO users (username, email, phone_number, full_name, surname) VALUES ($1, $2, $3, $4, $5) RETURNING id",
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
		"INSERT INTO customer (company_name, customer_type, user_id) VALUES ($1, $2, $3)",
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
