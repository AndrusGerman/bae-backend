package main

import (
	"fmt"

	"go.uber.org/fx"
)

func mainX() {
	type rutaString string
	type middlewareString string

	type HandlerAdd struct {
		Ruta       rutaString
		Middleware []middlewareString
	}

	var rutaModule = func(rutaName rutaString) fx.Option {
		return fx.Module(
			"ruta",
			fx.Decorate(func() rutaString {
				return rutaName // el decorate entero se sustituye por el constructor de handler
			}),
			fx.Provide(
				fx.Annotate(
					func(middlewares []middlewareString, ruta rutaString) *HandlerAdd {
						return &HandlerAdd{
							Ruta:       ruta,
							Middleware: middlewares,
						}
					},
					fx.ParamTags(`group:"middlewares"`),
					fx.ResultTags(`group:"handlersAdd"`),
				),
			),
		)
	}

	var middlewaresModule = func(middlewares []middlewareString) fx.Option {
		var middlewareProviders = make([]any, len(middlewares))
		for i := range middlewares {
			middlewareProviders[i] = fx.Annotate(
				func() middlewareString {
					return middlewares[i]
				},
				fx.ResultTags(`group:"middlewares"`),
			)
		}

		return fx.Module(
			"middleware",
			fx.Provide(
				middlewareProviders...,
			),
		)
	}

	var handlersModule = func(rutas []rutaString, middlewares []middlewareString) fx.Option {
		var modulesRoutes = make([]fx.Option, len(rutas))
		for i := range rutas {
			modulesRoutes[i] = rutaModule(rutas[i])
		}

		return fx.Module(
			"handlerModule",
			fx.Decorate(fx.Annotate(
				func() []middlewareString {
					return make([]middlewareString, 0)
				},
				fx.ResultTags(`group:"middlewares"`),
			)),
			middlewaresModule(middlewares),
			fx.Options(modulesRoutes...),
		)
	}

	fx.New(
		fx.Supply(rutaString(""), fx.Annotate(
			middlewareString(""),
			fx.ResultTags(`group:"middlewares"`),
		)),
		handlersModule(
			[]rutaString{"hello", "bye"},
			[]middlewareString{""},
		),
		handlersModule(
			[]rutaString{"secret"},
			[]middlewareString{"auth"},
		),
		fx.Invoke(
			fx.Annotate(
				func(handlers []*HandlerAdd) {
					for i := range handlers {
						fmt.Println("El handler ", handlers[i].Ruta, "tiene el middleware ", handlers[i].Middleware)
					}
				},
				fx.ParamTags(`group:"handlersAdd"`),
			),
		),
	).Run()

}
