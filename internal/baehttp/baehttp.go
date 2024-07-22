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

func (baeHttp *Bae) Use(middleware ...Middleware) *Bae {

	var middlewareGin []gin.HandlerFunc = make([]gin.HandlerFunc, len(middleware))
	for i := range middleware {
		middlewareGin[i] = baeToGinHandler(baeHttp, middleware[i].Handler())
	}
	baeHttp.core.Use(middlewareGin...)
	return baeHttp
}

func (baeHttp *Bae) NewContext(ctx *gin.Context) *Context {
	return NewContext(ctx, baeHttp)
}

func (baeHttp *Bae) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	return NewRouterGroup(baeHttp, relativePath, handlers...)
}

func (baeHttp *Bae) POST(relativePath string, handlers ...HandlerFunc) {
	baeHttp.core.POST(relativePath, baeHandlersToGin(baeHttp, handlers...)...)
}

func (baeHttp *Bae) Serve(listenAddr string) error {
	return baeHttp.core.Run(listenAddr)
}
