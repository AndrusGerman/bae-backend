package http

import (
	"bae-backend/internal/core/domain"
	"bae-backend/internal/core/port"

	"github.com/gin-gonic/gin"
)

// UserHandler represents the HTTP handler for user-related requests
type UserHandler struct {
	svc port.UserService
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(svc port.UserService) *UserHandler {
	return &UserHandler{
		svc,
	}
}

// registerRequest represents the request body for creating a user
type registerRequest struct {
	Name     string `json:"name" binding:"required" example:"John Doe"`
	Email    string `json:"email" binding:"required,email" example:"test@example.com"`
	Password string `json:"password" binding:"required,min=8" example:"12345678"`
}

// @Router			/users [post]
func (uh *UserHandler) Register(ctx *gin.Context) {
	var req registerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return
	}

	user := domain.User{
		// Name:     req.Name,
		// Email:    req.Email,
		// Password: req.Password,
	}

	_, err := uh.svc.Register(&user)
	if err != nil {
		//handleError(ctx, err)
		return
	}

	//handleSuccess(ctx, rsp)
}
