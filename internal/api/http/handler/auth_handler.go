package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mfturkcan/nereye-rest-api/internal/api/http/error_handler"
	"github.com/mfturkcan/nereye-rest-api/internal/api/http/server"
	"github.com/mfturkcan/nereye-rest-api/pkg/model"
	"github.com/mfturkcan/nereye-rest-api/pkg/repository"
)

type AuthHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
}

type CustomAuthHandler struct {
	userRepository *repository.CustomUserRepository
	logger         *log.Logger
}

func NewCustomAuthHandler(logger *log.Logger, userRepository *repository.CustomUserRepository, router *server.CustomRouter) *CustomUserHandler {
	handler := &CustomUserHandler{
		logger:         logger,
		userRepository: userRepository,
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

	user, err := h.userRepository.GetUserByUsernameAndPassword(dto.Username, dto.Password)

	if err != nil {
		error_handler.HandleDataCannotHandledError(w, r, h.logger)
		return
	}

	res := &model.UserLoginResponseDto{
		AccessToken:  user.Username,
		RefreshToken: "dümenden refresh token",
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&res)
}

func (h *CustomAuthHandler) RegisterRoutes(router *server.CustomRouter) {
	router.Router.Route("/auth", func(r chi.Router) {
		r.Get("/sign-in", h.Login)
	})
}
