package store

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

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
	var (
		Username         = os.Getenv("DB_USER")
		Password         = os.Getenv("DB_PASSWORD")
		Host             = os.Getenv("DB_HOST")
		Port             = os.Getenv("DB_PORT")
		DB               = os.Getenv("DB_NAME")
		Driver           = os.Getenv("DB_DRIVER")
		ConnectionString = fmt.Sprintf("%s://%s:%s@%s:%s:5432/%s?sslmode=disable", Driver, Username, Password, Host, Port, DB)
	)

	db, err := sql.Open(Driver, ConnectionString)

	if err != nil {
		store.logger.Println("Error creating instance of db")
		store.logger.Fatalln(err)
	}

	err = db.PingContext(*store.ctx)

	if err != nil {
		store.logger.Println("Error occured reaching db")
		store.logger.Println(err)
	}

	store.logger.Println("Successfully connected to db")
	store.DB = db
}

func (store *Store) autoMigrate() {
	//err := store.DB..AutoMigrate(&model.Todo{})

	// if err != nil {
	// 	store.logger.Panicln("Error occured while creating migrations")
	// }
}
