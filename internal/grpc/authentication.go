package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (s *Server) authValid(ctx context.Context, authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}

	applicationToken, err := s.db.ReadApplicationTokenByToken(ctx, authorization[0])
	if err != nil {
		logger.WithField("func", "authValid").Errorf("db read: %s", err.Error())
		return false
	}
	if applicationToken == nil {
		return false
	}

	return true
}

func (s *Server) unaryInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ErrMissingMetadata
	}
	fmt.Printf("md: %+v\n", md)

	if !s.authValid(ctx, md["authorization"]) {
		return nil, ErrInvalidToken
	}

	return handler(ctx, req)
}
