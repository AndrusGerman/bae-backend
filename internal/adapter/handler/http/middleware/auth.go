package middleware

import (
	"bae-backend/internal/baehttp"
	"fmt"

	"github.com/gin-gonic/gin"
)

func NewAuthMiddleware() baehttp.IMiddleware {
	return baehttp.NewGinMiddleware(func(ctx *gin.Context) {
		fmt.Println("Todo ok")

		ctx.Next()
	})
}
