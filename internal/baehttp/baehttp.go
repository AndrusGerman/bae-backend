package baehttp

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

type Bae struct {
	core           *gin.Engine
	errorStatusMap map[error]int
}

func NewBae() *Bae {
	var baeHttp = new(Bae)
	baeHttp.core = gin.New()
	baeHttp.errorStatusMap = make(map[error]int)

	return baeHttp
}

func (baeHttp *Bae) ErrorStatusMap(errorStatusMap map[error]int) *Bae {
	baeHttp.errorStatusMap = errorStatusMap
	return baeHttp
}

func (baeHttp *Bae) Use(middleware ...IMiddleware) *Bae {
	var middlewaresGin = make([]gin.HandlerFunc, len(middleware))
	for i := range middleware {
		middlewaresGin[i] = middleware[i].toGin()
	}
	baeHttp.core.Use(middlewaresGin...)
	return baeHttp
}

func (baeHttp *Bae) NewContextHandler(ctx *gin.Context) Context {
	return NewContextHandler(ctx, baeHttp)
}

func (baeHttp *Bae) Serve(listenAddr string) error {
	return baeHttp.core.Run(listenAddr)
}

func (baeHttp *Bae) AddHandlers(handlers ...Handler) *Bae {
	for i := range handlers {
		baeHttp.AddHandler(
			NewHandlerAdd(handlers[i]),
		)
	}
	return baeHttp
}

func (baeHttp *Bae) AddHandler(handlerAdd IHandlerAdd) *Bae {
	var handler = handlerAdd.GetHandler()
	var config = handler.Config()
	var middlewares = handlerAdd.GetMiddlewares()

	var ginHandler = baeHttp.newGinHandler(handler, middlewares)
	baeHttp.core.Handle(config.GetMethod(), config.GetPattern(), ginHandler)
	return baeHttp
}

func (baeHttp *Bae) newGinHandler(handler Handler, middlewares []IMiddleware) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err error
		for i := range middlewares {
			var baeContext = baeHttp.NewContextHandler(ctx)
			err = middlewares[i].Handler(baeContext)
			if err != nil {
				slog.Info("middleware error")
				return
			}
			if !baeContext.IsNext() {
				slog.Info("middleware is stop")
				return
			}
		}

		handler.Handler(baeHttp.NewContextHandler(ctx))
	}
}
