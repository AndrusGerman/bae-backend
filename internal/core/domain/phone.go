package domain

import "fmt"

type Phone struct {
	CallCode CallCode `json:"callCode" bson:"callCode"`
	Number   uint64   `json:"number" bson:"number"`
}

func (phone *Phone) FullPhone() string {
	return fmt.Sprintf("+%d%d", phone.CallCode.Int64(), phone.Number)
}
