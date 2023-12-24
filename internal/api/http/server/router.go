package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/mfturkcan/nereye-rest-api/internal/api/http/errors"
	m "github.com/mfturkcan/nereye-rest-api/internal/api/http/middleware"
)

type CustomRouter struct {
	logger *log.Logger
	Router *chi.Mux
}

func NewCustomRouter(logger *log.Logger) *CustomRouter {
	customRouter := &CustomRouter{
		logger: logger,
		Router: chi.NewRouter(),
	}

	customRouter.Setup()

	return customRouter
}

func (router *CustomRouter) Setup() {
	router.LoadCustomerMiddlewares()
	router.ConfigureCors()
	router.LoadCustomRoutes()
}

func (router *CustomRouter) ConfigureCors() {
	router.Router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"},
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
}

func (router *CustomRouter) LoadCustomerMiddlewares() {
	router.Router.Use(middleware.Logger)
	router.Router.Use(m.ContentTypeApplicationJsonMiddleware)
}

func (router *CustomRouter) LoadCustomRoutes() {
	router.Router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(&errors.ErrorResponse{
			Message: "Request not found",
			Status:  http.StatusNotFound,
			Detail:  "No such a request exists on path",
		})
	})

	router.Router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_ = json.NewEncoder(w).Encode(&errors.ErrorResponse{
			Message: "Method is not valid",
			Status:  http.StatusMethodNotAllowed,
		})
	})
}
