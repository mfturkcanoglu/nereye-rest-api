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

type CustomProductHandler struct {
	productRepository *repository.CustomProductRepository
	logger            *log.Logger
}

func NewCustomProductHandler(
	logger *log.Logger,
	productRepository *repository.CustomProductRepository,
	router *server.CustomRouter,
) *CustomProductHandler {
	handler := &CustomProductHandler{
		logger:            logger,
		productRepository: productRepository,
	}
	handler.RegisterRoutes(router)
	return handler
}

func (h *CustomProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := h.productRepository.GetAll()

	if err != nil {
		error_handler.HandleDataCannotHandledError(w, r, h.logger)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&res)
}

func (h *CustomProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	product := &model.ProductCreate{}
	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		error_handler.HandleInvalidSchemaError(w, r, err, h.logger)
		return
	}

	err = h.productRepository.Create(product)

	if err != nil {
		error_handler.HandleDataCannotHandledError(w, r, h.logger)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *CustomProductHandler) RegisterRoutes(router *server.CustomRouter) {
	router.Router.Route("/api/v1/product", func(r chi.Router) {
		r.Get("/", h.GetAll)
		r.Post("/", h.Create)
	})
}
