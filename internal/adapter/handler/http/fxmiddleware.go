package http

import (
	"bae-backend/internal/baehttp"

	"go.uber.org/fx"
)

func GlobalMiddlewaresModule(fxInject fx.Option, middleware ...any) fx.Option {
	var middlewareAnnotate = make([]any, len(middleware))
	for i := range middleware {
		middlewareAnnotate[i] = AsMiddleware(middleware[i], `group:"global_middleware"`)
	}

	return fx.Module("global_middleware", fxInject, fx.Provide(middlewareAnnotate...))
}

func AsMiddleware(f any, resultTag string) any {
	return fx.Annotate(
		f,
		fx.As(new(baehttp.IMiddleware)),
		fx.ResultTags(resultTag),
	)
}

func MiddlewareModule(resultTag string, middlewareConstructor ...any) fx.Option {
	var middlewareAnnotate = make([]any, len(middlewareConstructor))
	for i := range middlewareConstructor {
		middlewareAnnotate[i] = AsMiddleware(middlewareConstructor[i], resultTag)
	}
	return fx.Module("middleware", fx.Provide(middlewareAnnotate...))
}
