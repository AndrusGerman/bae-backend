package port

import "bae-backend/internal/core/domain"

//go:generate mockgen -source=message.go -destination=mock/message.go -package=mock

type MessageRepository interface {
	GetMessageByChat(chatId domain.Id) ([]domain.Message, error)
}

// MessageService is an interface for interacting with contact-related business logic
type MessageService interface {
	GetMessageByChat(chatId domain.Id) ([]domain.Message, error)
}
