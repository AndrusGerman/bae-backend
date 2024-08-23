package port

import "bae-backend/internal/core/domain"

//go:generate mockgen -source=chat.go -destination=mock/chat.go -package=mock

type ChatRepository interface {
	Create(contact *domain.Chat) (*domain.Chat, error)
	GetChats(userId domain.Id) ([]domain.Chat, error)
	FindGroups(name string) ([]domain.Chat, error)
}

// ContactService is an interface for interacting with contact-related business logic
type ChatService interface {
	Create(contact *domain.Chat) (*domain.Chat, error)
	GetChats(userId domain.Id) ([]domain.Chat, error)
	FindGroups(name string) ([]domain.Chat, error)
}
