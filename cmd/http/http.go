package main

import (
	"bae-backend/internal/adapter/handler/http"
	"bae-backend/internal/adapter/handler/http/hcountry"
	"bae-backend/internal/adapter/handler/http/huser"
	"bae-backend/internal/baehttp"
	"bae-backend/internal/core/domain"

	"go.uber.org/fx"
)

func baehttpModule() fx.Option {
	// global Middlewares
	var middlewares = http.GlobalMiddlewaresModule(
		fx.Supply(&baehttp.CorsConfig{AllowAllOrigins: true}),
		baehttp.Cors,
		baehttp.Recovery,
	)

	// routes
	var routes = http.RouterModule(
		http.NewRouterModuleConfig(),
		huser.NewUserGetAllHandler,
		huser.NewUserRegisterHandlerHandler,
		huser.NewUserGetUserByIdPhoneHandler,
		huser.NewUserGetUserByFullPhoneHandler,
		hcountry.NewCountryGetAllHandler,
		hcountry.NewCountryGetHandler,
	)

	// bae http module
	var baeHttp = fx.Module("baehttp",
		fx.Supply(baehttp.ErrorStatusMap(domain.ErrorStatusMap)),
		middlewares,
		routes,
		fx.Provide(
			http.NewHttpConfig,
			baehttp.NewBae,
		),
		// run http server
		fx.Invoke(http.RunHttpServer),
	)
	return baeHttp
}
