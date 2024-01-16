package model

type User struct {
	Username    string
	Email       string
	PhoneNumber string
	FullName    string
	Surname     string
	Verified    bool
	Enabled     bool
	Active      bool
	DefaultModel
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
