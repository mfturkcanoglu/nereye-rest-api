package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	m "github.com/mfturkcan/nereye-rest-api/internal/api/http/middleware"
)


func Setup() *chi.Mux{
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(m.ContentTypeApplicationJsonMiddleware)

	loadRoutes(r)
	loadCustomRoutes(r)

	return r
}

func loadRoutes(router *chi.Mux){
	router.Get("/", func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
		//w.Header().Add("Content-Type", "application/json")
		dict :=map[interface{}]interface{}{
		 "test":"api",
		}
		err := json.NewEncoder(w).Encode(dict)
		if err != nil {
			fmt.Println("Error", err)
		}
	})
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