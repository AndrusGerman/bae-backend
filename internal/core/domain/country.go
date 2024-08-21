package domain

import "github.com/biter777/countries"

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

func (country Country) Id() uint64 {
	return uint64(country)
}
