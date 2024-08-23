package port

import "bae-backend/internal/core/domain"

//go:generate mockgen -source=chatuser.go -destination=mock/chatuser.go -package=mock

type ChatUserRepository interface {
	Create(contact *domain.Chat) (*domain.Chat, error)
	GetChats(userId domain.Id) ([]domain.Chat, error)
}

// ChatUserService is an interface for interacting with contact-related business logic
type ChatUserService interface {
	Create(contact *domain.Chat) (*domain.Chat, error)
	GetChats(userId domain.Id) ([]domain.Chat, error)
}
