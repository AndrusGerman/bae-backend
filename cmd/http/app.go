package main

import (
	"bae-backend/internal/adapter/config"
	"bae-backend/internal/adapter/handler/http"
	"bae-backend/internal/adapter/handler/http/huser"
	"bae-backend/internal/adapter/storage/mongodb"
	"bae-backend/internal/adapter/storage/mongodb/repository"
	"bae-backend/internal/baehttp"
	"bae-backend/internal/core/service"
	"fmt"
	"log/slog"
	"os"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		//fx.NopLogger,
		fx.Provide(
			config.New,
			mongodb.New,
			repository.NewUserRepository,
			service.NewUserService,
			baehttp.NewBae,
			http.NewConfigureRouter,
		),
		http.RouterModule(
			huser.NewUserGetAllHandler,
			huser.NewUserRegisterHandlerHandler,
		),
		fx.Invoke(RunHttpServer),
	).Run()

}

func RunHttpServer(httpConfig *config.HTTP, baehttp *baehttp.Bae, routerConfiguration *http.RouterConfiguration) {
	// set base configuration
	baehttp.Use(routerConfiguration.Middleware...)
	baehttp.ErrorStatusMap(routerConfiguration.ErrorStatusMap)
	baehttp.AddHandlers(routerConfiguration.Handlers...)

	listenAddr := fmt.Sprintf("%s:%s", httpConfig.URL, httpConfig.Port)
	// Start server
	slog.Info("Starting the HTTP server", "listen_address", listenAddr)
	var err = baehttp.Serve(listenAddr)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}
