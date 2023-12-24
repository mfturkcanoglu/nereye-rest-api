package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/mfturkcan/nereye-rest-api/internal/api/http/server"
	"github.com/mfturkcan/nereye-rest-api/internal/store"
)

var (
	ctx    context.Context
	logger *log.Logger = log.Default()
	db     *sql.DB
)

func main() {
	ctx = context.Background()
	server := server.NewServer(logger, &ctx)

	store := store.NewStore(logger, &ctx)
	db = store.InitializeDatabase()

	server.ListenAndServe()
}
