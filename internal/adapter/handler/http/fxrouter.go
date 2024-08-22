package http

import (
	"bae-backend/internal/baehttp"
	"bae-backend/internal/core/util"

	"go.uber.org/fx"
)

func RouterModule(config *RouterModuleConfig, handlers ...any) fx.Option {
	var opts []fx.Option
	var middlewareTag = util.NewParamTag("middlewares", "group")
	opts = append(opts, fx.Module("router_config",
		MiddlewareModule(middlewareTag, config.MiddlewaresContructors...),
	))

	for i := range handlers {
		opts = append(opts, NewRouterAdd(handlers[i], middlewareTag))
	}

	return fx.Module(
		"routes",
		opts...,
	)
}

func NewRouterModuleConfig(middlewares ...any) *RouterModuleConfig {
	return &RouterModuleConfig{
		MiddlewaresContructors: middlewares,
	}
}

type RouterModuleConfig struct {
	MiddlewaresContructors []any
}

func NewRouterAdd(routerConstructor any, middlewareTag string) fx.Option {
	var handlerTag = util.NewParamTag("handler", "name")
	return fx.Module("routeAdd",
		fx.Provide(
			AsRoute(routerConstructor, handlerTag),
			fx.Annotate(
				func(handler baehttp.Handler, middlewares []baehttp.IMiddleware) baehttp.IHandlerAdd {
					return baehttp.NewHandlerAdd(handler, middlewares...)
				},
				fx.As(new(baehttp.IHandlerAdd)),
				fx.ResultTags(`group:"handlers_add"`),
				fx.ParamTags(handlerTag, middlewareTag),
			),
		),
	)
}

func AsRoute(routerConstructor any, paramTag string) any {
	return fx.Annotate(
		routerConstructor,
		fx.As(new(baehttp.Handler)),
		fx.ResultTags(paramTag),
	)
}
