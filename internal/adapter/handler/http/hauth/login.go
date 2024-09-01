package hauth

import (
	"bae-backend/internal/baehttp"
	"bae-backend/internal/core/domain"
	"bae-backend/internal/core/port"
	"fmt"
	"log"
	"net/http"
)

// UserHandler represents the HTTP handler for user-related requests
type LoginHandler struct {
	svc port.UserService
}

// NewUserHandler creates a new UserHandler instance
func NewLoginHandlerHandler(svc port.UserService) baehttp.Handler {
	return &LoginHandler{
		svc,
	}
}

func (uh *LoginHandler) Config() baehttp.HandlerConfig {
	return baehttp.NewHandlerConfig(http.MethodPost, base+"/login")
}

// registerRequest represents the request body for creating a user
type loginDtoRequest struct {
	Phone   string         `json:"phone" binding:"required"`
	Country domain.Country `json:"countryId" binding:"required"`
}

func (uh *LoginHandler) Handler(ctx baehttp.Context) error {
	var err error
	var req = new(loginDtoRequest)
	if err = ctx.BindJSON(req); err != nil {
		return ctx.HandleError(err)
	}

	var calcode = req.Country.CallCodes()[0]
	var fullPhone = fmt.Sprintf("%d%s", calcode.Int64(), req.Phone)

	log.Println("login with: ", fullPhone)

	rsp, err := uh.svc.GetByFullPhone(fullPhone)
	if err != nil {
		return ctx.HandleError(err)
	}

	return ctx.HandleSuccess(rsp)
}
