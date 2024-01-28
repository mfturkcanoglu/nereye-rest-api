package repository

import (
	"database/sql"
	"log"

	"github.com/mfturkcan/nereye-rest-api/pkg/model"
	"github.com/mfturkcan/nereye-rest-api/pkg/query"
)

type CustomerRepository interface {
	CreateCustomer(customerCreate *model.CustomerCreate) error
	GetAll() ([]*model.CustomerGet, error)
}

type CustomCustomerRepository struct {
	logger         *log.Logger
	db             *sql.DB
	userRepository *CustomUserRepository
}

func NewCustomerRepository(logger *log.Logger, db *sql.DB, userRepository *CustomUserRepository) *CustomCustomerRepository {
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
			&customer.ID,
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
	userId, err := repo.userRepository.CreateUser(&model.UserCreate{
		Username:    username,
		Email:       customerCreate.User.Email,
		PhoneNumber: customerCreate.User.PhoneNumber,
		FullName:    customerCreate.User.FullName,
		Surname:     customerCreate.User.Surname,
		Password:    customerCreate.User.Password,
	})

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
