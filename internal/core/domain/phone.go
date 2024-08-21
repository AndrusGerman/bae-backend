package domain

type Phone struct {
	Country  Country  `json:"country" bson:"countryId"`
	CallCode CallCode `json:"callCode" bson:"callCode"`
	Number   uint64   `json:"number" bson:"number"`
}
