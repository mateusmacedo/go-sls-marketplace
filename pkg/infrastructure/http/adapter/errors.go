package adapter

import (
	"errors"
)

var (
	ErrHttpMethodNotAllowed = errors.New("method not allowed")
	ErrHttpInvalidJSON      = errors.New("invalid JSON")
	ErrServiceError         = errors.New("some service error")
)
