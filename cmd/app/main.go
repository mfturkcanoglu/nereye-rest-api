package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/mfturkcan/nereye-rest-api/internal/api/http/handler"
	"github.com/mfturkcan/nereye-rest-api/internal/api/http/server"
	"github.com/mfturkcan/nereye-rest-api/internal/config"
	"github.com/mfturkcan/nereye-rest-api/internal/store"
	"github.com/mfturkcan/nereye-rest-api/pkg/repository"
)

var (
	ctx    context.Context = context.Background()
	logger *log.Logger     = log.Default()
)

func main() {
	var (
		store  *store.Store         = store.NewStore(logger, &ctx)
		router *server.CustomRouter = server.NewCustomRouter(logger)
		_      *config.Config       = config.NewConfig(logger).LoadEnv()
		db     *sql.DB              = store.InitializeDatabase()

		userRepository            *repository.CustomUserRepository            = repository.NewUserRepository(logger, db)
		customerRepository        *repository.CustomCustomerRepository        = repository.NewCustomerRepository(logger, db)
		restaurantRepository      *repository.CustomRestaurantRepository      = repository.NewRestaurantRepository(logger, db)
		restaurantPhotoRepository *repository.CustomRestaurantPhotoRepository = repository.NewRestaurantPhotoRepository(logger, db)
		categoryRepository        *repository.CustomCategoryRepository        = repository.NewCategoryRepository(logger, db)
		productRepository         *repository.CustomProductRepository         = repository.NewProductRepository(logger, db)
		_                         *handler.CustomUserHandler                  = handler.NewCustomUserHandler(logger, userRepository, router)
		_                         *handler.CustomCustomerHandler              = handler.NewCustomCustomerHandler(logger, customerRepository, router)
		_                         *handler.CustomRestaurantHandler            = handler.NewCustomRestaurantHandler(logger, restaurantRepository, restaurantPhotoRepository, router)
		_                         *handler.CustomCategoryHandler              = handler.NewCustomCategoryHandler(logger, categoryRepository, router)
		_                         *handler.CustomProductHandler               = handler.NewCustomProductHandler(logger, productRepository, router)
	)
	defer store.Close()

	server := server.NewServer(logger, &ctx, router)

	server.ListenAndServe()
}
