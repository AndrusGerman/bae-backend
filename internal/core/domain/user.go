package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ObjectID primitive.ObjectID

type User struct {
	Id      ObjectID `json:"Id" bson:"_id"`
	Country string   `json:"country" bson:"country"`
	Phone   string   `json:"phone" bson:"phone"`
	Email   string   `json:"email" bson:"email"`
}
