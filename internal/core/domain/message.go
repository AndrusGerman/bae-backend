package domain

import "time"

type MessageType uint

const (
	MessageTypeText     MessageType = 1
	MessageTypeImagen   MessageType = 2
	MessageTypeSticker  MessageType = 3
	MessageTypeVideo    MessageType = 4
	MessageTypeStatus   MessageType = 5
	MessageTypeCallInfo MessageType = 6
)

type MessageContent struct {
	Text string // TEXTO:
	Uri  string // URL -> sticker, IMAGEN,VIDEO
}

type Message struct {
	Id Id `json:"Id" bson:"_id"`

	FromUserId       Id             `json:"fromUserId" bson:"fromUserId"`
	Content          MessageContent `json:"content" bson:"content"`
	ChatId           Id             `json:"chatId" bson:"chatId"`
	ForwardMessageId Id             `json:"forwardMessageId" bson:"forwardMessageId"`
	MessageType      MessageType    `json:"messageTypeId" bson:"messageTypeId"`
	DateCreated      time.Time      `json:"dateCreated" bson:"dateCreated"`
}
