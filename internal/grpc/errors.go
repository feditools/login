package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// ErrMissingMetadata is returned when a request has invalid metadata.
	ErrMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	// ErrInvalidToken is returned when a request has an invalid token.
	ErrInvalidToken = status.Errorf(codes.Unauthenticated, "invalid token")
)
