package domain

import (
	"encoding"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//import "go.mongodb.org/mongo-driver/bson/primitive"

var _ encoding.TextMarshaler = Id{}
var _ encoding.TextUnmarshaler = &Id{}

type Id struct {
	value primitive.ObjectID
}

func (id Id) MarshalJSON() ([]byte, error) {
	return id.value.MarshalJSON()
}
func (id Id) MarshalText() ([]byte, error) {
	return id.value.MarshalText()
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

func (id *Id) UnmarshalJSON(b []byte) error {
	return id.value.UnmarshalJSON(b)
}

func (id *Id) UnmarshalText(b []byte) error {
	return id.value.UnmarshalText(b)
}
