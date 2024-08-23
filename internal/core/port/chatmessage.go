package port

import "bae-backend/internal/core/domain"

//go:generate mockgen -source=message.go -destination=mock/chatmessage.go -package=mock

type ChatMessageRepository interface {
	GetMessages(chatId domain.Id) ([]domain.Message, error)
}

type ChatMessageService interface {
	GetMessages(chatId domain.Id) ([]domain.Message, error)
}
