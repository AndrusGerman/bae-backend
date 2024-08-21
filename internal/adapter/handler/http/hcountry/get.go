package hcountry

import (
	"bae-backend/internal/baehttp"
	"bae-backend/internal/core/domain"
	"net/http"
)

// CountryGetHandler represents the HTTP handler for user-related requests
type CountryGetHandler struct {
}

// NewCountryGetHandler creates a new CountryGetHandler instance
func NewCountryGetHandler() baehttp.Handler {
	return &CountryGetHandler{}
}

func (uh *CountryGetHandler) Config() baehttp.HandlerConfig {
	return baehttp.NewHandlerConfig(http.MethodGet, base+"/:countryId")
}

func (uh *CountryGetHandler) Handler(ctx baehttp.Context) error {
	var countryId, err = ctx.Param("countryId").Uint64()
	if err != nil {
		return ctx.HandleError(err)
	}

	return ctx.HandleSuccess(domain.Country(countryId))
}
