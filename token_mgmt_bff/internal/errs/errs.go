package errs

import "errors"

// Common app or domain errors.
var (
	ErrTokenInvalid = errors.New("token is invalid")
	ErrTokenExpired = errors.New("token is expired")
)
