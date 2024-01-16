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

type CustomCategoryHandler struct {
	categoryRepository *repository.CustomCategoryRepository
	logger             *log.Logger
}

func NewCustomCategoryHandler(logger *log.Logger, categoryRepository *repository.CustomCategoryRepository, router *server.CustomRouter) *CustomCategoryHandler {
	handler := &CustomCategoryHandler{
		logger:             logger,
		categoryRepository: categoryRepository,
	}
	handler.RegisterRoutes(router)
	return handler
}

func (h *CustomCategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := h.categoryRepository.GetAll()

	if err != nil {
		msg := errors.HandleDataCannotHandledError(w, r)
		h.logger.Println(msg)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&res)
}

func (h *CustomCategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	category := &model.CategoryCreate{}
	err := json.NewDecoder(r.Body).Decode(&category)

	if err != nil {
		msg := errors.HandleInvalidSchemaError(w, r, err)
		h.logger.Println(msg)
		return
	}

	err = h.categoryRepository.Create(category)

	if err != nil {
		msg := errors.HandleDataCannotHandledError(w, r)
		h.logger.Println(msg)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *CustomCategoryHandler) RegisterRoutes(router *server.CustomRouter) {
	router.Router.Route("/api/v1/category", func(r chi.Router) {
		r.Get("/", h.GetAll)
		r.Post("/", h.Create)
	})
}
