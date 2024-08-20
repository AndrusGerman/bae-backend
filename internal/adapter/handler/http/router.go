package http

import (
	"bae-backend/internal/baehttp"
	"bae-backend/internal/core/domain"

	"go.uber.org/fx"
)

type RouterDto struct {
	fx.In
	Handlers []baehttp.Handler `group:"routes"`
	Bae      *baehttp.Bae
}

func DecorateHandlerConfiguration(routerDto RouterDto) *baehttp.Bae {
	// set base configuration
	routerDto.Bae.Use(
		baehttp.Cors(baehttp.CorsConfig{AllowAllOrigins: true}),
		baehttp.Recovery(),
	)
	routerDto.Bae.ErrorStatusMap(domain.ErrorStatusMap)
	routerDto.Bae.AddHandlers(routerDto.Handlers...)
	return routerDto.Bae
}

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(baehttp.Handler)),
		fx.ResultTags(`group:"routes"`),
	)
}

func RouterModule(handlers ...any) fx.Option {
	var routeAnnotate = make([]any, len(handlers))
	for i := range handlers {
		routeAnnotate[i] = AsRoute(handlers[i])
	}

	return fx.Module("routes", fx.Provide(routeAnnotate...))
}
