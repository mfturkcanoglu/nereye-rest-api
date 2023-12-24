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
	ctx    context.Context = context.Background()
	logger *log.Logger     = log.Default()
)

func main() {
	var (
		store          *store.Store                     = store.NewStore(logger, &ctx)
		db             *sql.DB                          = store.InitializeDatabase()
		router         *server.CustomRouter             = server.NewCustomRouter(logger)
		userRepository *repository.CustomUserRepository = repository.NewUserRepository(logger, db)
		_              *handler.CustomUserHandler       = handler.NewCustomerUserHandler(logger, userRepository, router)
	)

	defer store.Close()

	server := server.NewServer(logger, &ctx, router)

	server.ListenAndServe()
}
