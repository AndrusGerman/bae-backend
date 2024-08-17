package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Id primitive.ObjectID

func (id Id) MarshalJSON() ([]byte, error) {
	return primitive.ObjectID(id).MarshalJSON()
}
func (id Id) MarshalText() ([]byte, error) {
	return primitive.ObjectID(id).MarshalText()
}

func (id Id) Hex() string {
	return primitive.ObjectID(id).Hex()
}

func (id Id) UnmarshalText(b []byte) error {
	var x = primitive.NewObjectID()
	x.UnmarshalJSON(b)

	return nil
}
