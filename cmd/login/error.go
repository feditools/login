package main

import "strings"

// NewApplicationError wraps a message in an ApplicationError object.
func NewApplicationError(msg ...string) error {
	return &ApplicationError{message: strings.Join(msg, ": ")}
}

// ApplicationError is returned when the application throws an error.
type ApplicationError struct {
	message string
}

// Error returns the error message as a string.
func (e *ApplicationError) Error() string {
	return e.message
}
