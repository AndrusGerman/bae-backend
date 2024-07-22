package baehttp

import "github.com/gin-gonic/gin"

type HandlerFunc func(*Context)

func ginToBaeHandler(ginHandler gin.HandlerFunc) HandlerFunc {
	return func(ctx *Context) {
		ginHandler(ctx.ginCtx)
	}
}

func baeToGinHandler(baeHttp *Bae, baeHandler HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		baeHandler(baeHttp.NewContext(ctx))
	}
}

func baeHandlersToGin(baeHttp *Bae, baeHandler ...HandlerFunc) []gin.HandlerFunc {
	var handlers = make([]gin.HandlerFunc, len(baeHandler))

	for i := range handlers {
		handlers[i] = baeToGinHandler(baeHttp, baeHandler[i])
	}
	return handlers
}
