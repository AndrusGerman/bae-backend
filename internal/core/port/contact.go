package port

import "bae-backend/internal/core/domain"

//go:generate mockgen -source=contact.go -destination=mock/contact.go -package=mock

type ContactRepository interface {
	Create(contact *domain.Contact) (*domain.Contact, error)
	GetContacts(userId domain.Id) ([]domain.Contact, error)
}

// ContactService is an interface for interacting with contact-related business logic
type ContactService interface {
	Create(contact *domain.Contact) (*domain.Contact, error)
	GetContacts(userId domain.Id) ([]domain.Contact, error)
}
