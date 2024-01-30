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
	Username        string `json:"username"`
	Email           string `json:"email"`
	PhoneNumber     string `json:"phone_number"`
	FullName        string `json:"full_name"`
	Surname         string `json:"surname"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type UserLoginRequestDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginQueryResult struct {
	UserId       string
	PasswordHash string
}

type UserLoginResponseDto struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
