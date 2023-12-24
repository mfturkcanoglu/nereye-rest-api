package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mfturkcan/nereye-rest-api/internal/api/http/errors"
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

func (h *CustomUserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	res, err := h.userRepository.GetAll()

	if err != nil {
		msg := "Users cannot handled from db"
		h.logger.Println(msg)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(&errors.ErrorResponse{
			Message: msg,
			Status:  http.StatusInternalServerError,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&res)
}

func (h *CustomUserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &model.UserCreate{}
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		msg := "Schema is not correct"
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&errors.ErrorResponse{
			Message: msg,
			Detail:  err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	_, err = h.userRepository.CreateUser(user)

	if err != nil {
		msg := "Users cannot handled from db"
		h.logger.Println(msg)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(&errors.ErrorResponse{
			Message: msg,
			Status:  http.StatusInternalServerError,
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *CustomUserHandler) RegisterRoutes(router *server.CustomRouter) {
	router.Router.Route("/api/v1/user", func(r chi.Router) {
		r.Get("/", h.GetAllUsers)
		r.Post("/", h.CreateUser)
	})
}
