package domain

import (
	"github.com/biter777/countries"
)

type CallCode countries.CallCode
type CallCodes []CallCode

func (cd CallCode) Int64() int64 {
	return int64(cd)
}

func (cds CallCodes) In(calcode CallCode) bool {
	for _, cc := range cds {
		if int64(cc) == int64(calcode) {
			return true
		}
	}
	return false
}
