package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mfturkcan/nereye-rest-api/internal/api/http/error_handler"
	"github.com/mfturkcan/nereye-rest-api/internal/api/http/server"
	"github.com/mfturkcan/nereye-rest-api/internal/service"
	"github.com/mfturkcan/nereye-rest-api/pkg/model"
	"github.com/mfturkcan/nereye-rest-api/pkg/repository"
)

type AuthHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
}

type CustomAuthHandler struct {
	userRepository *repository.CustomUserRepository
	logger         *log.Logger
	authService    *service.AuthService
}

func NewCustomAuthHandler(logger *log.Logger, userRepository *repository.CustomUserRepository, router *server.CustomRouter, authService *service.AuthService) *CustomAuthHandler {
	handler := &CustomAuthHandler{
		logger:         logger,
		userRepository: userRepository,
		authService:    authService,
	}
	handler.RegisterRoutes(router)
	return handler
}

func (h *CustomAuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	dto := &model.UserLoginRequestDto{}
	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		error_handler.HandleInvalidSchemaError(w, r, err, h.logger)
		return
	}

	tokens, err := h.authService.Login(dto)

	if err != nil {
		error_handler.HandleDataCannotHandledError(w, r, h.logger)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&tokens)
}

func (h *CustomAuthHandler) RegisterRoutes(router *server.CustomRouter) {
	router.Router.Route("/auth", func(r chi.Router) {
		r.Post("/sign-in", h.Login)
	})
}
