package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (s *Server) unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	metric := s.metrics.NewGRPCRequest(info.FullMethod)
	l := logger.WithField("func", "unaryInterceptor")

	// get metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		go metric.Done(true)

		return nil, ErrMissingMetadata
	}

	// validate login
	if !s.authValid(ctx, md["authorization"]) {
		go metric.Done(true)

		return nil, ErrInvalidToken
	}

	// do request
	i, err := handler(ctx, req)
	if err != nil {
		errStatus := status.Convert(err)
		l.Warnf("grpc err: %s", errStatus.Code().String())

		go metric.Done(true)

		return nil, err
	}

	go metric.Done(false)

	return i, err
}
