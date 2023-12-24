package model

import (
	"fmt"
	"time"

	"math"

	"github.com/google/uuid"
)

type Customer struct {
	ID           uuid.UUID
	CompanyName  string
	CustomerType string
	AddressId    uuid.UUID
	UserId       uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
	Deleted      bool
}

type CustomerCreate struct {
	CompanyName     string `json:"company_name"`
	CustomerType    string `json:"customer_type"`
	Email           string `json:"email"`
	PhoneNumber     string `json:"phone_number"`
	FullName        string `json:"full_name"`
	Surname         string `json:"surname"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

func (customerCreate *CustomerCreate) CreateRandomCustomerUsername() string {
	maxLengthOfName := int(math.Min(float64(len(customerCreate.FullName)), 10))
	fullName := customerCreate.FullName[0:maxLengthOfName]

	maxLengthOfSurname := int(math.Min(float64(len(customerCreate.Surname)), 10))
	surname := customerCreate.FullName[0:maxLengthOfSurname]

	currentTime := time.Now().Unix()

	return fmt.Sprintf("%s%s%d", fullName, surname, currentTime)
}