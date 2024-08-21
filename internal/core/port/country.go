package port

import "bae-backend/internal/core/domain"

type CountryService interface {
	GetById(id uint64) (domain.Country, error)
	GetAll() []domain.Country
}
