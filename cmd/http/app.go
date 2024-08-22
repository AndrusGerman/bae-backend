package main

import (
	"bae-backend/internal/adapter/config"
	"bae-backend/internal/adapter/handler/http"
	"bae-backend/internal/adapter/handler/http/hcountry"
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
	var coreApp = fx.Provide(
		config.New,
		mongodb.New,
		repository.NewUserRepository,
		service.NewUserService,
		service.NewCountryService,
		baehttp.NewBae,
	)

	var globalHttpMiddleware = http.GlobalMiddlewaresModule(
		fx.Supply(&baehttp.CorsConfig{AllowAllOrigins: true}),
		baehttp.Cors,
		baehttp.Recovery,
	)

	fx.New(
		fx.NopLogger,
		coreApp,
		globalHttpMiddleware,
		http.RouterModule(
			http.NewRouterModuleConfig(),
			huser.NewUserGetAllHandler,
			huser.NewUserRegisterHandlerHandler,
			huser.NewUserGetUserByIdPhoneHandler,
			huser.NewUserGetUserByFullPhoneHandler,
			hcountry.NewCountryGetAllHandler,
			hcountry.NewCountryGetHandler,
		),
		fx.Decorate(http.DecorateBaeInject),
		fx.Invoke(RunHttpServer),
	).Run()

}

func RunHttpServer(httpConfig *config.HTTP, baehttp *baehttp.Bae) {

	listenAddr := fmt.Sprintf("%s:%s", httpConfig.URL, httpConfig.Port)
	// Start server
	slog.Info("Starting the HTTP server", "listen_address", listenAddr)
	var err = baehttp.Serve(listenAddr)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}
