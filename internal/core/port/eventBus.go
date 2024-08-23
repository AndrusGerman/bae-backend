package port

import "bae-backend/internal/core/domain"

// EventBusService is an interface for interacting with contact-related business logic
type EventBusService interface {
	Send(eventName domain.Event, data any)
	Subscribe(eventName domain.Event, calback domain.EventCalback)
}
