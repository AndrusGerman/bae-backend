package baehttp

import "strconv"

type Param string

func (param Param) Uint64() (uint64, error) {
	return strconv.ParseUint(string(param), 10, 64)
}

func (param Param) String() string {
	return string(param)
}
