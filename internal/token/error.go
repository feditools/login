package token

import "errors"

const (
	errRespGenerateToken = "couldn't generate token for %s: %s"
)

var (
	// ErrInvalidLength is returned when a token's data is an invalid length.
	ErrInvalidLength = errors.New("invalid length")
	// ErrInvalidTokenKind is returned when a token is an unexpected kind.
	ErrInvalidTokenKind = errors.New("invalid token kind")
	// ErrSaltEmpty is returned when a token's data is an invalid length.
	ErrSaltEmpty = errors.New("salt empty")
)
