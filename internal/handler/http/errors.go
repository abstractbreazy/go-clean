package http

import (
	"errors"
	"net/http"
)

type HTTPError struct {
	Code    int
	Message string
}

func NewHTTPError(code int, message string) HTTPError {
	return HTTPError{
		Code:    code,
		Message: message,
	}
}

type HTTPMessage string

const (
	BadRequestMsg         HTTPMessage = "Bad Request"
	ValidationMsg         HTTPMessage = "Validation error"
	SomethingWentWrongMsg HTTPMessage = "Something went wrong"
	TooManyRequestMsg     HTTPMessage = "Too many requests"
)

// Application errors.
var (
	ErrBadRequest    = errors.New("EBADREQUEST")
	ErrInternal      = errors.New("EINTERNAL")
	ErrUnprocessable = errors.New("EUNPROCESSABLE")
	ErrManyRequests  = errors.New("EMANYREQUESTS")
)

func ErrToHTTPStatus(err error) (int, HTTPMessage) {
	switch {
	case errors.Is(err, ErrBadRequest):
		return http.StatusBadRequest, BadRequestMsg
	case errors.Is(err, ErrUnprocessable):
		return http.StatusUnprocessableEntity, ValidationMsg
	case errors.Is(err, ErrManyRequests):
		return http.StatusTooManyRequests, TooManyRequestMsg
	case errors.Is(err, ErrInternal):
		return http.StatusInternalServerError, SomethingWentWrongMsg
	default:
		return 0, ""
	}
}
