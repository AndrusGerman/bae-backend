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
		fx.NopLogger,

		fx.Provide(
			config.New,
			mongodb.New,
			repository.NewUserRepository,
			service.NewUserService,
			baehttp.NewBae,
			http.AsRoute(huser.NewUserGetAllHandler),
			http.AsRoute(huser.NewUserRegisterHandlerHandler),
		),
		fx.Invoke(http.ConfigureRouter),
		fx.Invoke(RunHttpServer),
	).Run()

}

func RunHttpServer(config *config.Container, baehttp *baehttp.Bae) {
	listenAddr := fmt.Sprintf("%s:%s", config.HTTP.URL, config.HTTP.Port)
	// Start server
	slog.Info("Starting the HTTP server", "listen_address", listenAddr)
	var err = baehttp.Serve(listenAddr)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}
