package domain

type User struct {
	Id      Id     `json:"Id" bson:"_id"`
	Country string `json:"country" bson:"country"`
	Phone   string `json:"phone" bson:"phone"`
	Email   string `json:"email" bson:"email"`
}
