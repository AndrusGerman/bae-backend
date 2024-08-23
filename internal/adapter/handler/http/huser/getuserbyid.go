package huser

import (
	"bae-backend/internal/baehttp"
	"bae-backend/internal/core/domain"
	"bae-backend/internal/core/port"
	"net/http"
)

// UserGetUserByIdPhoneHandler represents the HTTP handler for user-related requests
type UserGetUserByIdPhoneHandler struct {
	userService port.UserService
}

// NewGetUserByIdPhoneHandler creates a new UserGetUserByIdPhoneHandler instance
func NewGetUserByIdPhoneHandler(userService port.UserService) baehttp.Handler {
	return &UserGetUserByIdPhoneHandler{
		userService: userService,
	}
}

func (uh *UserGetUserByIdPhoneHandler) Config() baehttp.HandlerConfig {
	return baehttp.NewHandlerConfig(http.MethodGet, base+"/:userId")
}

func (uh *UserGetUserByIdPhoneHandler) Handler(ctx baehttp.Context) error {
	var userId domain.Id
	var err error

	if userId, err = domain.NewIdFromHex(ctx.Param("userId").String()); err != nil {
		return ctx.HandleError(err)
	}

	rsp, err := uh.userService.GetUser(userId)
	if err != nil {
		return ctx.HandleError(err)
	}

	return ctx.HandleSuccess(rsp)
}
