package service

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	logger *log.Logger
	salt   int
}

func NewAuthService(logger *log.Logger, salt int) *AuthService {
	return &AuthService{
		logger: logger,
		salt:   14,
	}
}

func (service *AuthService) GenerateHashPassword(password string) (string, error) {
	salt := service.salt
	if salt == 0 {
		salt = 14
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(bytes), err
}

func (service *AuthService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
