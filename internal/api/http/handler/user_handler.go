package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/mfturkcan/nereye-rest-api/internal/api/http/error_handler"
	"github.com/mfturkcan/nereye-rest-api/internal/api/http/server"
	"github.com/mfturkcan/nereye-rest-api/pkg/model"
	"github.com/mfturkcan/nereye-rest-api/pkg/repository"
)

type UserHandler interface {
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
}

type CustomUserHandler struct {
	userRepository *repository.CustomUserRepository
	logger         *log.Logger
}

func NewCustomUserHandler(logger *log.Logger, userRepository *repository.CustomUserRepository, router *server.CustomRouter) *CustomUserHandler {
	handler := &CustomUserHandler{
		logger:         logger,
		userRepository: userRepository,
	}
	handler.RegisterRoutes(router)
	return handler
}

func (h *CustomUserHandler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	if strings.TrimSpace(username) == "" {
		error_handler.HandleInvalidSchemaError(w, r, errors.New("Invalid query"), h.logger)
		return
	}

	res, err := h.userRepository.GetUser(username)

	if err != nil {
		error_handler.HandleDataCannotHandledError(w, r, h.logger)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&res)
}

func (h *CustomUserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	res, err := h.userRepository.GetAll()

	if err != nil {
		error_handler.HandleDataCannotHandledError(w, r, h.logger)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&res)
}

func (h *CustomUserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &model.UserCreate{}
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		error_handler.HandleInvalidSchemaError(w, r, err, h.logger)
		return
	}

	_, err = h.userRepository.CreateUser(user)

	if err != nil {
		error_handler.HandleDataCannotHandledError(w, r, h.logger)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *CustomUserHandler) RegisterRoutes(router *server.CustomRouter) {
	router.Router.Route("/api/v1/user", func(r chi.Router) {
		r.Get("/all", h.GetAllUsers)
		r.Get("/", h.GetUserInfo)
		r.Post("/", h.CreateUser)
	})
}
