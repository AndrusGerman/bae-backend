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

func (baeHttp *Bae) AddHandlers(handlers ...IHandlerAdd) *Bae {
	for i := range handlers {
		baeHttp.AddHandler(
			handlers[i],
		)
	}
	return baeHttp
}

func (baeHttp *Bae) AddHandler(handlerAdd IHandlerAdd) *Bae {
	var handler = handlerAdd.GetHandler()
	var config = handler.Config()
	var middlewares = handlerAdd.GetMiddlewares()

	slog.Info("add: "+config.GetPattern(), "middlewarecount", len(middlewares))

	var ginHandler = baeHttp.newGinHandler(handler, middlewares)
	baeHttp.core.Handle(config.GetMethod(), config.GetPattern(), ginHandler)
	return baeHttp
}

func (baeHttp *Bae) newGinHandler(handler Handler, middlewares []IMiddleware) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		shouldReturn := baeHttp.runMiddlewares(middlewares, ctx)
		if shouldReturn {
			return
		}
		handler.Handler(baeHttp.NewContextHandler(ctx))
	}
}

func (baeHttp *Bae) runMiddlewares(middlewares []IMiddleware, ctx *gin.Context) bool {
	for i := range middlewares {
		var baeContext = baeHttp.NewContextHandler(ctx)
		var err = middlewares[i].Handler(baeContext)
		if err != nil {
			slog.Info("middleware error")
			return true
		}
		if !baeContext.IsNext() {
			slog.Info("middleware is stop")
			return true
		}
	}
	return false
}
