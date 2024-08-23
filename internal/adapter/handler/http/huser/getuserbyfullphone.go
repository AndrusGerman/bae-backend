package huser

import (
	"bae-backend/internal/baehttp"
	"bae-backend/internal/core/port"
	"net/http"
)

// UserGetUserByFullPhoneHandler represents the HTTP handler for user-related requests
type UserGetUserByFullPhoneHandler struct {
	userService port.UserService
}

// NewGetUserByFullPhoneHandler creates a new UserGetUserByFullPhoneHandler instance
func NewGetUserByFullPhoneHandler(userService port.UserService) baehttp.Handler {
	return &UserGetUserByFullPhoneHandler{
		userService: userService,
	}
}

func (uh *UserGetUserByFullPhoneHandler) Config() baehttp.HandlerConfig {
	return baehttp.NewHandlerConfig(http.MethodGet, base+"/by/fullphone/:fullPhone")
}

func (uh *UserGetUserByFullPhoneHandler) Handler(ctx baehttp.Context) error {
	var err error
	var fullPhone = ctx.Param("fullPhone").String()

	rsp, err := uh.userService.GetByFullPhone(fullPhone)
	if err != nil {
		return ctx.HandleError(err)

	}

	return ctx.HandleSuccess(rsp)
}
