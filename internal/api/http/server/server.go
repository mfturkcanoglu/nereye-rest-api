package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	router *CustomRouter
	logger *log.Logger
	ctx    *context.Context
}

func NewServer(log *log.Logger, ctx *context.Context, router *CustomRouter) *Server {
	s := &Server{
		router: router,
		logger: log,
		ctx:    ctx,
	}
	return s
}

func (s *Server) ListenAndServe() {
	httpServer := http.Server{
		Addr:         os.Getenv("APP_PORT"),
		Handler:      s.router.Router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  2 * time.Minute,
	}

	s.logger.Println("Starting server on", httpServer.Addr)

	s.logger.Fatalln(httpServer.ListenAndServe())
}
