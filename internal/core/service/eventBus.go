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

func (eb *EventBusService) Send(eventName domain.Event, data any) {
	for i := range eb.storeEvents[eventName] {
		eb.storeEvents[eventName][i](data)
	}
}

func (eb *EventBusService) Subscribe(eventName domain.Event, calback domain.EventCalback) {
	eb.storeEvents[eventName] = append(eb.storeEvents[eventName], calback)
}
