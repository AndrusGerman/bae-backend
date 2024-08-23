package domain

type Event string

const (
	EventChatDelete Event = "DeleteChat"
	EventChatCreate Event = "CreateChat"
)

type EventCalback func(data any)
