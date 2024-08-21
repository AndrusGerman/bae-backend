package service

import (
	"bae-backend/internal/core/domain"
	"bae-backend/internal/core/port"
	"errors"

	"github.com/biter777/countries"
)

type CountryService struct {
}

func NewCountryService() port.CountryService {
	return &CountryService{}
}

func (countryService *CountryService) GetById(id uint64) (domain.Country, error) {
	var country = domain.Country(id)

	if country.IsUnknown() {
		return 0, errors.New("is unknown country")
	}
	return country, nil
}

func (countryService *CountryService) GetAll() []domain.Country {
	var raw = countries.All()
	var resp = make([]domain.Country, len(raw))
	for i := range raw {
		resp[i] = domain.Country(raw[i])
	}
	return resp
}
