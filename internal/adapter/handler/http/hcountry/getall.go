package hcountry

import (
	"bae-backend/internal/baehttp"
	"bae-backend/internal/core/domain"
	"net/http"
)

// CountryGetAllHandler represents the HTTP handler for user-related requests
type CountryGetAllHandler struct {
}

// NewCountryGetAllHandler creates a new CountryGetAllHandler instance
func NewCountryGetAllHandler() baehttp.Handler {
	return &CountryGetAllHandler{}
}

func (uh *CountryGetAllHandler) Config() baehttp.HandlerConfig {
	return baehttp.NewHandlerConfig(http.MethodGet, base)
}

func (uh *CountryGetAllHandler) Handler(ctx baehttp.Context) error {
	return ctx.HandleSuccess(domain.ContryAll())
}
