package http

import (
	"bae-backend/internal/adapter/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Router is a wrapper for HTTP router
type Router struct {
	*gin.Engine
}

// NewRouter creates a new HTTP router
func NewRouter(
	config *config.HTTP,
	userHandler *UserHandler,
) (*Router, error) {
	// CORS

	router := gin.New()
	var configCors = cors.DefaultConfig()
	configCors.AllowAllOrigins = true
	router.Use(gin.Recovery(), cors.New(configCors))

	v1 := router.Group("/v1")
	{
		user := v1.Group("/users")
		{
			user.POST("/", userHandler.Register)
			// user.POST("/login", authHandler.Login)

			// authUser := user.Group("/")
			// {
			// 	authUser.GET("/", userHandler.ListUsers)
			// 	authUser.GET("/:id", userHandler.GetUser)

			// }
		}

	}

	return &Router{
		router,
	}, nil
}

// Serve starts the HTTP server
func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}
