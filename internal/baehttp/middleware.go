package baehttp

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Middleware interface {
	Handler() HandlerFunc
}

// Cors
type CorsConfig struct {
	AllowAllOrigins bool
}

func (ctx *CorsConfig) Handler() HandlerFunc {
	var configCors = cors.DefaultConfig()
	configCors.AllowAllOrigins = ctx.AllowAllOrigins
	return ginToBaeHandler(cors.New(configCors))
}

// Recovery
type Recovery struct {
}

func (ctx *Recovery) Handler() HandlerFunc {
	return ginToBaeHandler(gin.Recovery())
}
