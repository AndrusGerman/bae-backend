package baehttp

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Middleware interface {
	HandlerBase
	toGin() gin.HandlerFunc
	setBaeContext(bae *Bae)
}

func NewMiddleware(handlerBase HandlerBase) Middleware {
	return &HandlerMiddleware{base: handlerBase}
}

type HandlerMiddleware struct {
	base HandlerBase
	bae  *Bae
}

func (hm *HandlerMiddleware) setBaeContext(bae *Bae) {
	hm.bae = bae
}

func (hm *HandlerMiddleware) toGin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		hm.base.Handler(NewContextHandler(ctx, hm.bae))
	}
}

func (hm *HandlerMiddleware) Handler(ctx Context) error {
	return hm.base.Handler(ctx)
}

func NewGinMiddleware(ginMiddleware gin.HandlerFunc) Middleware {
	return &GinMiddleware{
		base: ginMiddleware,
	}
}

type GinMiddleware struct {
	base gin.HandlerFunc
}

func (gm *GinMiddleware) Handler(ctx Context) error {
	gm.base(ctx.getGin())
	if ctx.getGin().IsAborted() {
		return nil
	}
	ctx.Next()
	return nil
}

func (gm *GinMiddleware) toGin() gin.HandlerFunc {
	return gm.base
}

func (gm *GinMiddleware) setBaeContext(bae *Bae) {}

// Middlewares Globals
type CorsConfig struct {
	AllowAllOrigins bool
}

func Cors(config *CorsConfig) Middleware {
	var configCors = cors.DefaultConfig()
	configCors.AllowAllOrigins = config.AllowAllOrigins
	return NewGinMiddleware(cors.New(configCors))

}

func Recovery() Middleware {
	return NewGinMiddleware(gin.Recovery())
}
