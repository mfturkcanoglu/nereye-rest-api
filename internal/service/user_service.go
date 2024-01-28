package service

import (
	"errors"
	"log"

	"github.com/mfturkcan/nereye-rest-api/pkg/model"
	"github.com/mfturkcan/nereye-rest-api/pkg/repository"
)

type UserService struct {
	logger      *log.Logger
	userRepo    *repository.CustomUserRepository
	authService *AuthService
}

func NewUserService(logger *log.Logger, repo *repository.CustomUserRepository, authService *AuthService) *UserService {
	return &UserService{
		logger:      logger,
		userRepo:    repo,
		authService: authService,
	}
}

func (service *UserService) CreateUser(userCreate *model.UserCreate) error {
	if userCreate.Password != userCreate.PasswordConfirm {
		return errors.New("passwords does not match")
	}

	hashedPass, err := service.authService.GenerateHashPassword(userCreate.Password)

	if err != nil {
		return errors.New("error during creating password")
	}

	userCreate.Password = hashedPass

	_, err = service.userRepo.CreateUser(userCreate)
	return err
}
