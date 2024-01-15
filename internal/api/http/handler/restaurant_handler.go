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

type RestaurantHandler interface {
	GetCustomerRestaurants(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	CreateRestaurant(w http.ResponseWriter, r *http.Request)
}

type CustomRestaurantHandler struct {
	restaurantRepository      *repository.CustomRestaurantRepository
	restaurantPhotoRepository *repository.CustomRestaurantPhotoRepository
	logger                    *log.Logger
}

func NewCustomRestaurantHandler(logger *log.Logger,
	restaurantRepository *repository.CustomRestaurantRepository,
	resturantPhotoRepository *repository.CustomRestaurantPhotoRepository,
	router *server.CustomRouter) *CustomRestaurantHandler {
	handler := &CustomRestaurantHandler{
		logger:                    logger,
		restaurantRepository:      restaurantRepository,
		restaurantPhotoRepository: resturantPhotoRepository,
	}
	handler.RegisterRoutes(router)
	return handler
}

func (h *CustomRestaurantHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	customerId := r.URL.Query().Get("customer-id")
	res, err := h.restaurantRepository.GetAll(customerId)

	if err != nil {
		msg := "Restaurants cannot handled from db"
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

func (h *CustomRestaurantHandler) GetRestaurantPhotos(w http.ResponseWriter, r *http.Request) {
	restaurantId := r.URL.Query().Get("restaurant-id")

	res, err := h.restaurantPhotoRepository.GetAll(restaurantId)

	if err != nil {
		msg := "Restaurant photos cannot handled from db"
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

func (h *CustomRestaurantHandler) CreateRestaurant(w http.ResponseWriter, r *http.Request) {
	restaurant := &model.RestaurantCreate{}
	err := json.NewDecoder(r.Body).Decode(&restaurant)

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

	err = h.restaurantRepository.CreateRestaurant(restaurant)

	if err != nil {
		msg := "Restaurants cannot handled from db"
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

func (h *CustomRestaurantHandler) RegisterRoutes(router *server.CustomRouter) {
	router.Router.Route("/api/v1/restaurant", func(r chi.Router) {
		r.Get("/", h.GetAll)
		r.Post("/", h.CreateRestaurant)
		r.Get("/photos", h.GetRestaurantPhotos)
	})
}
