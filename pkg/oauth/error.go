package oauth

import "fmt"

// CallbackError is returned when the callback handler encounters an error.
type CallbackError struct {
	Code    int
	Message string
}

// Error returns the error message as a string.
func (e *CallbackError) Error() string {
	return fmt.Sprintf("%s (%d)", e.Message, e.Code)
}

// NewCallbackError wraps a message in a CallbackError object.
func NewCallbackError(c int, m string) error {
	return &CallbackError{
		Code:    c,
		Message: m,
	}
}
