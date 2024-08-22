package baehttp

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type IMiddleware interface {
	HandlerBase
	toGin() gin.HandlerFunc
}

func NewGinMiddleware(ginMiddleware gin.HandlerFunc) IMiddleware {
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

// Cors
type CorsConfig struct {
	AllowAllOrigins bool
}

func Cors(config *CorsConfig) IMiddleware {
	var configCors = cors.DefaultConfig()
	configCors.AllowAllOrigins = config.AllowAllOrigins
	return NewGinMiddleware(cors.New(configCors))

}

func Recovery() IMiddleware {
	return NewGinMiddleware(gin.Recovery())
}
