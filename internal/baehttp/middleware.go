package baehttp

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type IMiddleware interface {
	getGinMiddleware() gin.HandlerFunc
}

func NewGinMiddleware(ginMiddleware gin.HandlerFunc) IMiddleware {
	return &GinMiddleware{
		ginMiddleware: ginMiddleware,
	}
}

type GinMiddleware struct {
	ginMiddleware gin.HandlerFunc
}

func (ctx *GinMiddleware) getGinMiddleware() gin.HandlerFunc {
	return ctx.ginMiddleware
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
