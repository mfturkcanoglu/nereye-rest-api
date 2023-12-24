package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	m "github.com/mfturkcan/nereye-rest-api/internal/api/http/middleware"
)

var (
	logger *log.Logger = log.Default()
)


func (s *Server) Setup(){
	s.router.Use(middleware.Logger)
	s.router.Use(m.ContentTypeApplicationJsonMiddleware)

	loadRoutes(s.router)
	loadCustomRoutes(s.router)
}

func loadRoutes(router *chi.Mux){
	
}

func loadCustomRoutes(router *chi.Mux){
	type customRouteResponse struct {
		Message string `json:"message"`
		Status int `json:"status"`
	}

	router.NotFound(func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(&customRouteResponse{
			Message:"Request not found" ,
			Status: http.StatusNotFound,
		})
	})

	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusMethodNotAllowed)
		_ = json.NewEncoder(w).Encode(&customRouteResponse{
			Message:"Method is not valid" ,
			Status: http.StatusMethodNotAllowed,
		})
	})
}