package domain

import (
	"github.com/biter777/countries"
)

type CallCode countries.CallCode
type CallCodes []CallCode

func (cd CallCode) Int64() int64 {
	return int64(cd)
}
