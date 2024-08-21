package domain

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
)

type PhoneBase struct {
	Country  Country  `json:"countryId" bson:"countryId"`
	CallCode CallCode `json:"callCode" bson:"callCode"`
	Number   uint64   `json:"number" bson:"number"`
}

type Phone struct {
	PhoneBase
}

func (phone Phone) MarshalJSON() ([]byte, error) {
	type PhoneJSONResponse struct {
		PhoneBase
		Alpha string `json:"alpha"`
	}
	return json.Marshal(PhoneJSONResponse{
		PhoneBase: phone.PhoneBase,
		Alpha:     phone.Country.Alpha(),
	})
}

func (phone *Phone) UnmarshalJSON(b []byte) error {
	var err = json.Unmarshal(b, &phone.PhoneBase)
	if err != nil {
		return err
	}

	if !phone.Country.CallCodes().In(phone.CallCode) {
		return ErrThiCallCodeIsNotFound
	}
	return nil
}

func (phone Phone) MarshalBSON() ([]byte, error) {
	return bson.Marshal(phone.PhoneBase)
}

func (phone *Phone) UnmarshalBSON(b []byte) error {
	return bson.Unmarshal(b, &phone.PhoneBase)
}
