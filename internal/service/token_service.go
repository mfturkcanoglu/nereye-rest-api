package service

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenService struct {
	logger      *log.Logger
	TokenSecret string
}

type Claims struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

func NewTokenService(logger *log.Logger, tokenSecret string) *TokenService {
	return &TokenService{
		logger:      logger,
		TokenSecret: tokenSecret,
	}
}

func (service *TokenService) CreateAccessToken(userId string) (string, error) {
	var now time.Time = time.Now()
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(time.Minute * 15).Unix(),
		},
	})

	return accessToken.SignedString([]byte(service.TokenSecret))
}

func (service *TokenService) CreateRefreshToken() (string, error) {
	var now time.Time = time.Now()

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(time.Hour * 48).Unix(),
	})

	return refreshToken.SignedString([]byte(service.TokenSecret))
}

func (service *TokenService) ParseAccessToken(accessToken string) *Claims {
	parsedAccessToken, _ := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(service.TokenSecret), nil
	})

	return parsedAccessToken.Claims.(*Claims)
}

func (service *TokenService) ParseRefreshToken(refreshToken string) *jwt.StandardClaims {
	parsedRefreshToken, _ := jwt.ParseWithClaims(refreshToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(service.TokenSecret), nil
	})

	return parsedRefreshToken.Claims.(*jwt.StandardClaims)
}

func (service *TokenService) IsTokenValid(claims jwt.StandardClaims) bool {
	return claims.Valid() == nil
}
