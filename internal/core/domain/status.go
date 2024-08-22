package domain

type Status struct {
	Id        Id `json:"Id" bson:"_id"`
	MessageId Id `json:"messageId" bson:"messageId"`
}
