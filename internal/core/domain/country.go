package domain

import (
	"encoding/json"

	"github.com/biter777/countries"
)

type CountryInfo struct {
	Alpha     string `json:"alpha"`
	CountryId uint   `json:"countryId"`
	Name      string `json:"name"`
	Emoji     string `json:"emoji"`
}

type Country countries.CountryCode

func ContryAll() []Country {
	var raw = countries.All()
	var resp = make([]Country, len(raw))
	for i := range raw {
		resp[i] = Country(raw[i])
	}
	return resp
}

func (country Country) Alpha() string {
	return countries.CountryCode(country).Alpha2()
}

func (country Country) CallCodes() CallCodes {
	var callcodes = countries.CountryCode(country).CallCodes()
	var resp = make(CallCodes, len(callcodes))
	for i := range callcodes {
		resp[i] = CallCode(callcodes[i])
	}
	return resp
}

func (country Country) String() string {
	return countries.CountryCode(country).String()
}

func (country Country) Emoji() string {
	return countries.CountryCode(country).Emoji3()
}

func (country Country) Id() uint64 {
	return uint64(country)
}

func (country Country) Info() *CountryInfo {
	return &CountryInfo{
		Alpha:     country.Alpha(),
		CountryId: uint(country.Id()),
		Name:      country.String(),
		Emoji:     country.Emoji(),
	}
}

func (country Country) MarshalJSON() ([]byte, error) {
	return json.Marshal(country.Info())
}
