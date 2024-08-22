package http

import (
	"bae-backend/internal/baehttp"
	"bae-backend/internal/core/domain"
	"fmt"

	"go.uber.org/fx"
)

type RouterDto struct {
	fx.In
	Handlers   []baehttp.Handler     `group:"routes"`
	Middleware []baehttp.IMiddleware `group:"global_middleware"`
	Bae        *baehttp.Bae
}

func DecorateHandlerConfiguration(dto RouterDto) *baehttp.Bae {
	return dto.Bae.
		// set middlewares
		Use(dto.Middleware...).
		// set response error status map
		ErrorStatusMap(domain.ErrorStatusMap).
		// add handlers
		AddHandlers(dto.Handlers...)
}

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(baehttp.Handler)),
		fx.ResultTags(`group:"routes"`),
	)
}

func RouterModule(fxInject fx.Option, handlers ...any) fx.Option {
	var routeAnnotate = make([]any, len(handlers))
	for i := range handlers {
		routeAnnotate[i] = AsRoute(handlers[i])
	}

	return fx.Module(
		"routes",
		fxInject,
		fx.Decorate(DecorateInjectInternalMiddleware),
		fx.Provide(routeAnnotate...),
	)
}

func GlobalMiddlewaresModule(fxInject fx.Option, middleware ...any) fx.Option {
	var middlewareAnnotate = make([]any, len(middleware))
	for i := range middleware {
		middlewareAnnotate[i] = AsGlobalMiddleware(middleware[i])
	}

	return fx.Module("global_middleware", fxInject, fx.Provide(middlewareAnnotate...))
}

func AsGlobalMiddleware(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(baehttp.IMiddleware)),
		fx.ResultTags(`group:"global_middleware"`),
	)
}

func AsInternalMiddleware(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(baehttp.IMiddleware)),
		fx.ResultTags(`group:"internal_middleware"`),
	)
}

func MiddlewareModule(middleware ...any) fx.Option {
	var middlewareAnnotate = make([]any, len(middleware))
	for i := range middleware {
		middlewareAnnotate[i] = AsInternalMiddleware(middleware[i])
	}
	return fx.Module("middleware", fx.Provide(middlewareAnnotate...))
}

type DecorateInjectInternalMiddlewareResp struct {
	fx.Out
	Handlers []baehttp.Handler `group:"routes"`
}

type DecorateInjectInternalMiddlewareDto struct {
	fx.In
	Middleware []baehttp.IMiddleware `group:"internal_middleware"`
	Handlers   []baehttp.Handler     `group:"routes"`
}

func DecorateInjectInternalMiddleware(handlers []baehttp.Handler, middleware []baehttp.IMiddleware) DecorateInjectInternalMiddlewareResp {
	fmt.Println("El decorador se ejecuto")
	for i := range handlers {
		var config = handlers[i].Config()
		fmt.Println("En la ruta: ", config.GetPattern(), " se seteo middlewares ", len(middleware))
		//handlers[i].Config().SetMiddleware(middleware)
	}

	return DecorateInjectInternalMiddlewareResp{
		Handlers: handlers,
	}
}
