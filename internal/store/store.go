package store

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Store struct {
	DB     *sql.DB
	logger *log.Logger
	ctx    *context.Context
}

func NewStore(logger *log.Logger, ctx *context.Context) *Store {
	store := &Store{
		DB:     nil,
		logger: logger,
		ctx:    ctx,
	}
	return store
}

func (store *Store) InitializeDatabase() *sql.DB {
	store.createConnection()
	store.autoMigrate()
	return store.DB
}

func (store *Store) Close() {
	err := store.DB.Close()
	if err != nil {
		store.logger.Panicln("Error occured while closing connection")
	}
}

func (store *Store) createConnection() {
	const connectionString = "postgres://mfturkcan:@localhost:5432/nereyedb?sslmode=disable"

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		store.logger.Println("Error creating instance of db")
		store.logger.Fatalln(err)
	}

	err = db.PingContext(*store.ctx)

	if err != nil {
		store.logger.Println("Error occured reaching db")
		store.logger.Println(err)
	}

	store.DB = db
}

func (store *Store) autoMigrate() {
	//err := store.DB..AutoMigrate(&model.Todo{})

	// if err != nil {
	// 	store.logger.Panicln("Error occured while creating migrations")
	// }
}
