package http

import (
	"bae-backend/internal/baehttp"
	"bae-backend/internal/core/domain"

	"go.uber.org/fx"
)

type RouterDto struct {
	fx.In
	Handlers []baehttp.Handler `group:"routes"`
}

type RouterConfiguration struct {
	Handlers       []baehttp.Handler
	Middleware     []baehttp.IMiddleware
	ErrorStatusMap map[error]int
}

// ConfigureRouter creates a new HTTP router
func NewConfigureRouter(rdto RouterDto) *RouterConfiguration {
	return &RouterConfiguration{
		Handlers: rdto.Handlers,
		Middleware: []baehttp.IMiddleware{
			baehttp.Cors(baehttp.CorsConfig{AllowAllOrigins: true}),
			baehttp.Recovery(),
		},
		ErrorStatusMap: domain.ErrorStatusMap,
	}
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
