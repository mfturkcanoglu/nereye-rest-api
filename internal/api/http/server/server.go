package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router *chi.Mux
	logger *log.Logger
	ctx    *context.Context
}

func NewServer(log *log.Logger, ctx *context.Context) *Server {
	s := &Server{
		router: chi.NewRouter(),
		logger: logger,
		ctx:    ctx,
	}
	s.Setup()
	return s
}

func (s *Server) ListenAndServe() {
	httpServer := http.Server{
		Addr:         ":3000",
		Handler:      s.router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  2 * time.Minute,
	}

	s.logger.Fatalln(httpServer.ListenAndServe())
}
