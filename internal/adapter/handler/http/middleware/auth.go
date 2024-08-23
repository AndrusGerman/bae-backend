package middleware

import (
	"bae-backend/internal/baehttp"
)

func NewAuthMiddleware() baehttp.Middleware {
	return baehttp.NewMiddleware(&AuthMiddleware{})
}

type AuthMiddleware struct {
}

func (authMiddleware *AuthMiddleware) Handler(ctx baehttp.Context) error {

	ctx.Next()
	return nil
}
