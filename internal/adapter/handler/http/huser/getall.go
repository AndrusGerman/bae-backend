package huser

import (
	"bae-backend/internal/baehttp"
	"bae-backend/internal/core/port"
	"net/http"
)

// UserGetAllHandler represents the HTTP handler for user-related requests
type UserGetAllHandler struct {
	userService port.UserService
}

// NewUserGetAllHandler creates a new UserGetAllHandler instance
func NewUserGetAllHandler(userService port.UserService) baehttp.Handler {
	return &UserGetAllHandler{
		userService: userService,
	}
}

func (uh *UserGetAllHandler) Config() baehttp.HandlerConfig {
	return baehttp.NewHandlerConfig(http.MethodGet, base)

}

func (uh *UserGetAllHandler) Handler(ctx baehttp.Context) error {
	rsp, err := uh.userService.GetAllUsers()
	if err != nil {
		return ctx.HandleError(err)

	}

	return ctx.HandleSuccess(rsp)
}
