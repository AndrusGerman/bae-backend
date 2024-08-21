package domain

import (
	"errors"
	"net/http"
)

var (
	// ErrInternal is an error for when an internal service fails to process the request
	ErrInternal = errors.New("internal error")

	ErrThisElementIsAlredyExist = errors.New("this element is alredy exist")

	// ErrInternal is not implement
	ErrNotImplement = errors.New("err not implement")

	ErrThiCallCodeIsNotFound = errors.New("this callcode is not found in this country")
)

var ErrorStatusMap = map[error]int{
	ErrInternal:     http.StatusInternalServerError,
	ErrNotImplement: http.StatusNotFound,
}
