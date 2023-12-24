package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID
	Username    string
	Email       string
	PhoneNumber string
	FullName    string
	Surname     string
	Verified    bool
	Enabled     bool
	Active      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	Deleted     bool
}

type UserGet struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	FullName    string `json:"full_name"`
	Surname     string `json:"surname"`
}

type UserCreate struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	FullName    string `json:"full_name"`
	Surname     string `json:"surname"`
}
