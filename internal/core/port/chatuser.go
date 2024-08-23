package port

import "bae-backend/internal/core/domain"

//go:generate mockgen -source=chatuser.go -destination=mock/chatuser.go -package=mock

type ChatUserRepository interface {
	GetChats(userId domain.Id) ([]domain.ChatWithUser, error)
}

// ChatUserService is an interface for interacting with contact-related business logic
type ChatUserService interface {
	GetChats(userId domain.Id) ([]domain.ChatWithUser, error)
}
