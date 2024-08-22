package middleware

import (
	"bae-backend/internal/baehttp"

	"github.com/gin-gonic/gin"
)

func NewAuthMiddleware() baehttp.IMiddleware {
	return baehttp.NewGinMiddleware(func(ctx *gin.Context) {
		ctx.Next()
	})
}
