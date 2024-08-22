package domain

type BlockUser struct {
	Id Id `json:"Id" bson:"_id"`

	ToUserId   Id `json:"toUserId" bson:"toUserId"`
	FromUserId Id `json:"fromUserId" bson:"fromUserId"`
}
