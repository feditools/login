package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
)

func authValid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")

	return token == "some-secret-token"
}

func unaryInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ErrMissingMetadata
	}

	if !authValid(md["authorization"]) {
		return nil, ErrInvalidToken
	}

	return handler(ctx, req)
}
