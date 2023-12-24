package main

import (
	"log"

	"github.com/mfturkcan/nereye-rest-api/internal/api/http/server"
	"github.com/mfturkcan/nereye-rest-api/internal/store"
	"gorm.io/gorm"
)

var (
	logger *log.Logger  = log.Default()
	db *gorm.DB
)

func main(){
	server := server.NewServer(logger)

	store := store.NewStore(logger)
	db = store.InitializeDatabase()

	server.ListenAndServe()
}