package domain

type Contact struct {
	Id Id `json:"Id" bson:"_id"`

	ToUserId   Id `json:"toUserId" bson:"toUserId"`
	FromUserId Id `json:"fromUserId" bson:"fromUserId"`

	ContactName string `json:"contactName" bson:"contactName"`
}
