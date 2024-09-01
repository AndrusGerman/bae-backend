package port

import "bae-backend/internal/core/domain"

// EventBusService is an interface for interacting with contact-related business logic
type EventBusService interface {
	Publish(eventName domain.Event, data any) error
	Subscribe(eventType domain.Event, handler domain.EventCalback) error
}
