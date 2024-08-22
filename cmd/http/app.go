package main

import (
	"bae-backend/internal/adapter/config"
	"bae-backend/internal/adapter/storage/mongodb"
	"bae-backend/internal/adapter/storage/mongodb/repository"
	"bae-backend/internal/core/service"

	"go.uber.org/fx"
)

func main() {
	// core modules
	var coreApp = fx.Provide(
		config.New,
		mongodb.New,
		repository.NewUserRepository,
		service.NewUserService,
		service.NewCountryService,
	)
	var baeHttp = baehttpModule()

	fx.New(
		//fx.NopLogger,
		coreApp,
		baeHttp,
	).Run()

}
