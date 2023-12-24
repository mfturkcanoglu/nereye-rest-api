package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mfturkcan/nereye-rest-api/internal/api/http/errors"
	"github.com/mfturkcan/nereye-rest-api/pkg/model"
	"github.com/mfturkcan/nereye-rest-api/pkg/repository"
)

type UserHandler struct {
	userRepository repository.UserRepository
	logger         *log.Logger
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
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

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
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

	err = h.userRepository.CreateUser(user)

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
