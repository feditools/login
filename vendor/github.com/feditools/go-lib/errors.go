package lib

import "errors"

var (
	// ErrInvalidAccountFormat is returned when a federated account is in an invalid format
	ErrInvalidAccountFormat = errors.New("invalid account format")
)
