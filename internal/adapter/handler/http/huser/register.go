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

func (uh *UserRegisterHandler) Config() baehttp.HandlerConfig {
	return baehttp.NewHandlerConfig(http.MethodPost, base+"/register")
}

// registerRequest represents the request body for creating a user
type registerDtoRequest struct {
	Phone    domain.Phone   `json:"phone" binding:"required"`
	UserName string         `json:"userName" binding:"required"`
	Country  domain.Country `json:"countryId" binding:"required"`
}

func (uh *UserRegisterHandler) Handler(ctx baehttp.Context) error {
	var err error
	var req = new(registerDtoRequest)
	if err = ctx.BindJSON(req); err != nil {
		return ctx.HandleError(err)
	}

	user := domain.User{
		Phone:    req.Phone,
		UserName: req.UserName,
		Country:  req.Country,
	}

	rsp, err := uh.svc.Register(&user)
	if err != nil {
		return ctx.HandleError(err)
	}

	return ctx.HandleSuccess(rsp)
}
