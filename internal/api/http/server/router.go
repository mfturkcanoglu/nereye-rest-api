package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	router.Router.Use(middleware.Logger)
	router.Router.Use(m.ContentTypeApplicationJsonMiddleware)

	loadCustomRoutes(router.Router)
}

func loadCustomRoutes(router *chi.Mux) {
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(&errors.ErrorResponse{
			Message: "Request not found",
			Status:  http.StatusNotFound,
			Detail:  "No such a request exists on path",
		})
	})

	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_ = json.NewEncoder(w).Encode(&errors.ErrorResponse{
			Message: "Method is not valid",
			Status:  http.StatusMethodNotAllowed,
		})
	})
}
