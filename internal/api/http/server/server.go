package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Server struct{
	router *chi.Mux
	logger *log.Logger
}

func NewServer(log *log.Logger) *Server {
	s := &Server{
		router: chi.NewRouter(),
		logger: logger,
	}
	s.Setup()
	return s
}

func (s *Server) ListenAndServe(){
	httpServer := http.Server{
		Addr:         ":3000",
		Handler:      s.router,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		IdleTimeout:  2 * time.Minute,
	}

	s.logger.Fatalln(httpServer.ListenAndServe())
}