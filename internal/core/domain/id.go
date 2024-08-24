package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Id struct {
	value primitive.ObjectID
}

func NewId() Id {
	return Id{
		value: primitive.NewObjectID(),
	}
}

func NewIdFromHex(id string) (Id, error) {
	var value, err = primitive.ObjectIDFromHex(id)
	return Id{value: value}, err
}

func (id Id) Hex() string {
	return id.value.Hex()
}

func (id Id) String() string {
	return id.value.String()
}

func (id Id) IsZero() bool {
	return id.value.IsZero()
}
func (id Id) Timestamp() time.Time {
	return id.value.Timestamp()
}

func (id Id) MarshalJSON() ([]byte, error) {
	return id.value.MarshalJSON()
}

func (id Id) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(id.value)
}

func (id Id) MarshalText() ([]byte, error) {
	return id.value.MarshalText()
}

func (id *Id) UnmarshalJSON(b []byte) error {
	return id.value.UnmarshalJSON(b)
}

func (id *Id) UnmarshalBSONValue(t bsontype.Type, b []byte) error {
	return bson.UnmarshalValue(t, b, &id.value)
}

func (id *Id) UnmarshalText(b []byte) error {
	return id.value.UnmarshalText(b)
}
