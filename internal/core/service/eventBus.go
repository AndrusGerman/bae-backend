package service

import (
	"bae-backend/internal/core/domain"
	"bae-backend/internal/core/port"
)

type EventBusService struct {
	storeEvents map[domain.Event][]domain.EventCalback
}

func NewEventBusService() port.EventBusService {
	return &EventBusService{}
}

func (eb *EventBusService) Publish(eventName domain.Event, data any) error {
	for i := range eb.storeEvents[eventName] {
		eb.storeEvents[eventName][i](data)
	}
	return nil
}

func (eb *EventBusService) Subscribe(eventName domain.Event, handler domain.EventCalback) error {
	eb.storeEvents[eventName] = append(eb.storeEvents[eventName], handler)
	return nil
}
