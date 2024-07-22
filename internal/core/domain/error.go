package domain

import (
	"errors"
	"net/http"
)

var (
	// ErrInternal is an error for when an internal service fails to process the request
	ErrInternal = errors.New("internal error")

	// ErrInternal is not implement
	ErrNotImplement = errors.New("err not implement")
)

var ErrorStatusMap = map[error]int{
	ErrInternal:     http.StatusInternalServerError,
	ErrNotImplement: http.StatusNotFound,
}
