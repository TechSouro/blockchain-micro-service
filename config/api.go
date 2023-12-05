package config

import (
	handler "service/internal/ApiModule"
	repository "service/internal/ApiModule/repository"
	routes "service/internal/ApiModule/routes"
	usecase "service/internal/ApiModule/usecase"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeAPI(connectionString string) *mux.Router {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	apiRepo := repository.NewApiRepository(db)
	apiUseCase := usecase.NewApiUseCase(apiRepo)
	apiHandler := handler.NewApiHandler(apiUseCase)

	return routes.ConfigureAPIRoutes(apiHandler)
}
