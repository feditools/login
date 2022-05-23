package oauth

import "fmt"

// Error is an oauth specific error.
type Error struct {
	Code    int
	Message string
}

// Error returns the error message as a string.
func (e *Error) Error() string {
	return fmt.Sprintf("%s (%d)", e.Message, e.Code)
}

// NewError wraps a message in a Error object.
func NewError(c int, m string) error {
	return &Error{
		Code:    c,
		Message: m,
	}
}
