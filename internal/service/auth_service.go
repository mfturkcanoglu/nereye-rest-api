package service

import (
	"errors"
	"log"

	"github.com/mfturkcan/nereye-rest-api/pkg/model"
	"github.com/mfturkcan/nereye-rest-api/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	logger         *log.Logger
	salt           int
	userRepository *repository.CustomUserRepository
	tokenService   *TokenService
}

func NewAuthService(logger *log.Logger, salt int, userRepository *repository.CustomUserRepository, tokenService *TokenService) *AuthService {
	return &AuthService{
		logger:         logger,
		salt:           14,
		userRepository: userRepository,
		tokenService:   tokenService,
	}
}

func (service *AuthService) GenerateHashPassword(password string) (string, error) {
	salt := service.salt
	if salt == 0 {
		salt = 12
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(bytes), err
}

func (service *AuthService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (service *AuthService) Login(dto *model.UserLoginRequestDto) (*model.UserLoginResponseDto, error) {
	// Find user with username

	queryRes, err := service.userRepository.GetUserIdByUsername(dto.Username)

	if err != nil {
		return nil, err
	}

	// Check if password correct
	isMatch := service.CheckPasswordHash(dto.Password, queryRes.PasswordHash)

	if !isMatch {
		return nil, errors.New("user credentianls are not correct")
	}

	refreshToken, err := service.tokenService.CreateRefreshToken()

	if err != nil {
		return nil, err
	}

	accessToken, err := service.tokenService.CreateAccessToken(queryRes.UserId)

	if err != nil {
		return nil, err
	}

	return &model.UserLoginResponseDto{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (service *AuthService) SignUp(userCreate *model.UserCreate) (*model.UserLoginResponseDto, error) {
	if userCreate.Password != userCreate.PasswordConfirm {
		return nil, errors.New("passwords does not match")
	}

	hashedPass, err := service.GenerateHashPassword(userCreate.Password)

	if err != nil {
		return nil, errors.New("error during creating password")
	}

	userCreate.Password = hashedPass

	userId, err := service.userRepository.CreateUser(userCreate)

	if err != nil {
		return nil, err
	}

	refreshToken, err := service.tokenService.CreateRefreshToken()

	if err != nil {
		return nil, err
	}

	accessToken, err := service.tokenService.CreateAccessToken(userId)

	if err != nil {
		return nil, err
	}
	return &model.UserLoginResponseDto{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
