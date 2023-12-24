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

var (
	logger *log.Logger = log.Default()
)

func (s *Server) Setup() {
	s.router.Use(middleware.Logger)
	s.router.Use(m.ContentTypeApplicationJsonMiddleware)

	loadRoutes(s.router)
	loadCustomRoutes(s.router)
}

func loadRoutes(router *chi.Mux) {

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
