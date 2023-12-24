package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/mfturkcan/nereye-rest-api/internal/api/http/handler"
	"github.com/mfturkcan/nereye-rest-api/internal/api/http/server"
	"github.com/mfturkcan/nereye-rest-api/internal/store"
	"github.com/mfturkcan/nereye-rest-api/pkg/repository"
)

var (
	ctx            context.Context
	logger         *log.Logger          = log.Default()
	router         *server.CustomRouter = server.NewCustomRouter(logger)
	db             *sql.DB
	userRepository *repository.CustomUserRepository = repository.NewUserRepository(logger, db)
	userHandler    *handler.CustomUserHandler       = handler.NewCustomerUserHandler(logger, userRepository, router)
)

func main() {
	ctx = context.Background()

	server := server.NewServer(logger, &ctx, router)

	store := store.NewStore(logger, &ctx)
	db = store.InitializeDatabase()

	server.ListenAndServe()
}
