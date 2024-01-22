package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	CompanyName  string
	CustomerType string
	AddressId    uuid.UUID
	UserId       uuid.UUID
	DefaultModel
}

type CustomerGet struct {
	ID           string `json:"customer_id"`
	CompanyName  string `json:"company_name"`
	CustomerType string `json:"customer_type"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_number"`
	FullName     string `json:"full_name"`
	Surname      string `json:"surname"`
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
	maxLengthOfName := min(len(customerCreate.FullName), 5)
	fullName := customerCreate.FullName[0:maxLengthOfName]

	maxLengthOfSurname := min(len(customerCreate.Surname), 5)
	surname := customerCreate.Surname[0:maxLengthOfSurname]

	currentTime := time.Now().Unix()

	return fmt.Sprintf("%s%s%d", fullName, surname, currentTime)
}
