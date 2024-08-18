package huser

import (
	"bae-backend/internal/baehttp"
	"bae-backend/internal/core/domain"
	"bae-backend/internal/core/port"
	"net/http"
)

// UserHandler represents the HTTP handler for user-related requests
type UserRegisterHandler struct {
	svc port.UserService
}

// NewUserHandler creates a new UserHandler instance
func NewUserRegisterHandlerHandler(svc port.UserService) baehttp.Handler {
	return &UserRegisterHandler{
		svc,
	}
}

func (uh *UserRegisterHandler) Config() *baehttp.Config {
	return &baehttp.Config{
		Pattern: base + "/register",
		Method:  http.MethodPost,
	}
}

// registerRequest represents the request body for creating a user
type registerDtoRequest struct {
	Name     string `json:"name" binding:"required" example:"John Doe"`
	Email    string `json:"email" binding:"required,email" example:"test@example.com"`
	Password string `json:"password" binding:"required,min=8" example:"12345678"`
}

func (uh *UserRegisterHandler) Handler(ctx *baehttp.Context) error {
	var err error
	var req = new(registerDtoRequest)
	if err = ctx.BindJSON(req); err != nil {
		return ctx.HandleError(err)
	}

	user := domain.User{
		Email: req.Email,
	}

	rsp, err := uh.svc.Register(&user)
	if err != nil {
		return ctx.HandleError(err)
	}

	return ctx.HandleSuccess(rsp)
}
