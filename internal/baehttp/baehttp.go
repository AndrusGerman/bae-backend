package baehttp

import "github.com/gin-gonic/gin"

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
		middlewaresGin[i] = middleware[i].getGinMiddleware()
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

func (baeHttp *Bae) Add(handler Handler) *Bae {
	var config = handler.Config()
	baeHttp.core.Handle(config.GetMethod(), config.GetPattern(), func(ctx *gin.Context) {
		handler.Handler(baeHttp.NewContextHandler(ctx))
	})
	return baeHttp
}
