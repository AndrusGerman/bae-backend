package domain

import (
	"encoding"
	"log"
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
	log.Println("MarshalJSON")
	return id.value.MarshalJSON()
}
func (id Id) MarshalText() ([]byte, error) {
	log.Println("MarshalText")
	return id.value.MarshalText()
}

func (id Id) Hex() string {
	log.Println("Hex")
	return id.value.Hex()
}

func (id Id) String() string {
	log.Println("String")
	return id.value.String()
}

func (id Id) UnmarshalJSON(b []byte) error {
	log.Println("UnmarshalJSON")
	return id.value.UnmarshalJSON(b)
}

func (id Id) UnmarshalText(b []byte) error {
	log.Println("UnmarshalText")
	return id.value.UnmarshalText(b)
}
func (id Id) IsZero() bool {
	log.Println("IsZero")
	return id.value.IsZero()
}
func (id Id) Timestamp() time.Time {
	log.Println("Timestamp")
	return id.value.Timestamp()
}
