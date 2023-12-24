package repository

import (
	"database/sql"
	"log"

	"github.com/mfturkcan/nereye-rest-api/pkg/model"
	"github.com/mfturkcan/nereye-rest-api/pkg/query"
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

func (repo *CustomCustomerRepository) GetAll() ([]*model.CustomerGet, error) {
	rows, err := repo.db.Query(query.CustomerUserSelectQuery())

	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	defer rows.Close()

	customers := []*model.CustomerGet{}

	for rows.Next() {
		customer := &model.CustomerGet{}
		err := rows.Scan(
			&customer.CompanyName,
			&customer.CustomerType,
			&customer.Username,
			&customer.PhoneNumber,
			&customer.Email,
			&customer.FullName,
			&customer.Surname,
		)

		if err != nil {
			repo.logger.Println(err)
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (repo *CustomCustomerRepository) CreateCustomer(customerCreate *model.CustomerCreate) error {
	username := customerCreate.CreateRandomCustomerUsername()
	var userId string
	err := repo.db.QueryRow(
		query.UserInsertQuery(),
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
