package huser

import (
	"bae-backend/internal/baehttp"
	"bae-backend/internal/core/port"
	"net/http"
)

// UserGetAllHandler represents the HTTP handler for user-related requests
type UserGetAllHandler struct {
	svc port.UserService
}

// NewUserGetAllHandler creates a new UserGetAllHandler instance
func NewUserGetAllHandler(svc port.UserService) baehttp.Handler {
	return &UserGetAllHandler{
		svc: svc,
	}
}

func (uh *UserGetAllHandler) Config() baehttp.HandlerConfig {
	return baehttp.NewHandlerConfig(http.MethodGet, base)

}

func (uh *UserGetAllHandler) Handler(ctx baehttp.Context) error {
	rsp, err := uh.svc.GetAll()
	if err != nil {
		return ctx.HandleError(err)

	}

	return ctx.HandleSuccess(rsp)
}
