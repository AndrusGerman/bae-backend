package baehttp

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors
type CorsConfig struct {
	AllowAllOrigins bool
}

func Cors(config CorsConfig) HandlerFunc {
	var configCors = cors.DefaultConfig()
	configCors.AllowAllOrigins = config.AllowAllOrigins
	return ginToBaeHandler(cors.New(configCors))
}

func Recovery() HandlerFunc {
	return ginToBaeHandler(gin.Recovery())
}
