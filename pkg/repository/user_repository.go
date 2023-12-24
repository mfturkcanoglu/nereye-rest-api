package repository

import (
	"database/sql"
	"log"

	"github.com/mfturkcan/nereye-rest-api/pkg/model"
)

type UserRepository interface {
	CreateUser(userCreate *model.UserCreate) error
	GetAll() ([]*model.UserGet, error)
}

type userRepository struct {
	logger *log.Logger
	db     *sql.DB
}

func NewUserRepository(logger *log.Logger, db *sql.DB) *userRepository {
	repo := &userRepository{
		logger: logger,
		db:     db,
	}
	return repo
}

func (repo *userRepository) CreateUser(userCreate *model.UserCreate) error {
	_, err := repo.db.Exec(
		"INSERT INTO users (username, email, phone_number, full_name, surname) VALUES ($1, $2, $3, $4, $5)",
		userCreate.Username, userCreate.Email, userCreate.PhoneNumber, userCreate.FullName, userCreate.Surname,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) GetAll() ([]*model.UserGet, error) {
	rows, err := repo.db.Query(`
		SELECT
		u.username,
		u.phone,
		u.email,
		u.full_name,
		u.last_name
		from users u
		order by u.updated_at desc
	`)

	if err != nil {
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
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
