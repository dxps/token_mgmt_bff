package errs

import "errors"

// Common app or domain errors.
var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrTokenInvalid       = errors.New("token is invalid")
	ErrTokenExpired       = errors.New("token is expired")
	ErrNotFound           = errors.New("not found")
)
