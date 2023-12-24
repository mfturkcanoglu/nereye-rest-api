package store

import (
	"log"

	"github.com/mfturkcan/nereye-rest-api/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store struct{
	DB *gorm.DB
	logger *log.Logger
}

func NewStore(logger *log.Logger) *Store {
	store := &Store{
		DB: nil,
		logger: logger,
	}
	return store
}

func (store *Store) InitializeDatabase() *gorm.DB{
	store.createConnection()
	store.autoMigrate()
	return store.DB
}

func (store *Store) Close(){
}

func (store *Store) createConnection(){
	const connectionString = "postgres://mfturkcan:@localhost:5432/nereyedb"

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		store.logger.Println("Error occured while connecting db")
		store.logger.Fatalln(err)
	}

	store.DB = db
}

func (store *Store) autoMigrate(){
	err := store.DB.AutoMigrate(&model.Todo{})
	
	if err != nil {
		store.logger.Panicln("Error occured while creating migrations")
	}
}