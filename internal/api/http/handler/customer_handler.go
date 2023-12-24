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

type CustomerHandler interface {
	GetAllCustomer(w http.ResponseWriter, r *http.Request)
	CreateCustomer(w http.ResponseWriter, r *http.Request)
}

type CustomCustomerHandler struct {
	customerRepository *repository.CustomCustomerRepository
	logger             *log.Logger
}

func NewCustomCustomerHandler(logger *log.Logger, customerRepository *repository.CustomCustomerRepository, router *server.CustomRouter) *CustomCustomerHandler {
	handler := &CustomCustomerHandler{
		logger:             logger,
		customerRepository: customerRepository,
	}
	handler.RegisterRoutes(router)
	return handler
}

func (h *CustomCustomerHandler) GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	res, err := h.customerRepository.GetAll()

	if err != nil {
		msg := "Customers cannot handled from db"
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

func (h *CustomCustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	customer := &model.CustomerCreate{}
	err := json.NewDecoder(r.Body).Decode(&customer)

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

	err = h.customerRepository.CreateCustomer(customer)

	if err != nil {
		msg := "Customers cannot handled from db"
		h.logger.Println(msg)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(&errors.ErrorResponse{
			Message: msg,
			Status:  http.StatusInternalServerError,
			Detail:  err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *CustomCustomerHandler) RegisterRoutes(router *server.CustomRouter) {
	router.Router.Route("/customer", func(r chi.Router) {
		r.Get("/", h.GetAllCustomer)
		r.Post("/add", h.CreateCustomer)
	})
}
