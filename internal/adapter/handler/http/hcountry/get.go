package hcountry

import (
	"bae-backend/internal/baehttp"
	"bae-backend/internal/core/domain"
	"bae-backend/internal/core/port"
	"net/http"
)

// CountryGetHandler represents the HTTP handler for user-related requests
type CountryGetHandler struct {
	countryService port.CountryService
}

// NewCountryGetHandler creates a new CountryGetHandler instance
func NewCountryGetHandler(countryService port.CountryService) baehttp.Handler {
	return &CountryGetHandler{
		countryService: countryService,
	}
}

func (uh *CountryGetHandler) Config() baehttp.HandlerConfig {
	return baehttp.NewHandlerConfig(http.MethodGet, base+"/:countryId")
}

func (uh *CountryGetHandler) Handler(ctx baehttp.Context) error {
	var country domain.Country
	var err error
	var countryId uint64

	countryId, err = ctx.Param("countryId").Uint64()
	if err != nil {
		return ctx.HandleError(err)
	}

	country, err = uh.countryService.GetById(countryId)
	if err != nil {
		return ctx.HandleError(err)
	}

	return ctx.HandleSuccess(country)
}
