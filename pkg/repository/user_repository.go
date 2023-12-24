package repository

import (
	"database/sql"
	"log"

	"github.com/mfturkcan/nereye-rest-api/pkg/model"
	"github.com/mfturkcan/nereye-rest-api/pkg/query"
)

type UserRepository interface {
	CreateUser(userCreate *model.UserCreate) error
	GetAll() ([]*model.UserGet, error)
}

type CustomUserRepository struct {
	logger *log.Logger
	db     *sql.DB
}

func NewUserRepository(logger *log.Logger, db *sql.DB) *CustomUserRepository {
	repo := &CustomUserRepository{
		logger: logger,
		db:     db,
	}
	return repo
}

func (repo *CustomUserRepository) CreateUser(userCreate *model.UserCreate) error {
	_, err := repo.db.Exec(
		query.UserInsertQuery(),
		userCreate.Username, userCreate.Email, userCreate.PhoneNumber, userCreate.FullName, userCreate.Surname,
	)
	if err != nil {
		repo.logger.Println(err)
		return err
	}
	return nil
}

func (repo *CustomUserRepository) GetAll() ([]*model.UserGet, error) {
	rows, err := repo.db.Query(query.UserSelectQuery())

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
