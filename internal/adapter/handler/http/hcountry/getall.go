package hcountry

import (
	"bae-backend/internal/baehttp"
	"bae-backend/internal/core/port"
	"net/http"
)

// CountryGetAllHandler represents the HTTP handler for user-related requests
type CountryGetAllHandler struct {
	countryService port.CountryService
}

// NewCountryGetAllHandler creates a new CountryGetAllHandler instance
func NewCountryGetAllHandler(
	countryService port.CountryService,
) baehttp.Handler {
	return &CountryGetAllHandler{
		countryService: countryService,
	}
}

func (uh *CountryGetAllHandler) Config() baehttp.HandlerConfig {
	return baehttp.NewHandlerConfig(http.MethodGet, base)
}

func (uh *CountryGetAllHandler) Handler(ctx baehttp.Context) error {
	return ctx.HandleSuccess(uh.countryService.GetAll())
}
