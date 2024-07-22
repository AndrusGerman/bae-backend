package http

import (
	"bae-backend/internal/adapter/config"
	"bae-backend/internal/adapter/handler/http/huser"
	"bae-backend/internal/baehttp"
	"bae-backend/internal/core/domain"
)

// NewRouter creates a new HTTP router
func NewRouter(
	config *config.HTTP,
	userHandler *huser.UserHandler,
) (*baehttp.Bae, error) {
	// CORS

	var coreBae = baehttp.NewBae().
		// add default middleware
		Use(
			baehttp.Cors(baehttp.CorsConfig{AllowAllOrigins: true}),
			baehttp.Recovery(),
		).
		// add status map erros
		ErrorStatusMap(domain.ErrorStatusMap)

	v1 := coreBae.Group("/v1")
	{
		user := v1.Group("/users")
		{
			user.GET("/", userHandler.All)
			user.POST("/", userHandler.Register)
			// user.POST("/login", authHandler.Login)

			// authUser := user.Group("/")
			// {
			// 	authUser.GET("/", userHandler.ListUsers)
			// 	authUser.GET("/:id", userHandler.GetUser)

			// }
		}

	}

	return coreBae, nil
}
