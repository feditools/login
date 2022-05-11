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
		go func() {
			ended := metric.Done(true)
			l.Debugf("rendering %s took %d ms", info.FullMethod, ended.Milliseconds())
		}()

		return nil, ErrMissingMetadata
	}

	// validate login
	if !s.authValid(ctx, md["authorization"]) {
		go func() {
			ended := metric.Done(true)
			l.Debugf("rendering %s took %d ms", info.FullMethod, ended.Milliseconds())
		}()

		return nil, ErrInvalidToken
	}

	// do request
	i, err := handler(ctx, req)
	if err != nil {
		errStatus := status.Convert(err)
		l.Warnf("grpc err: %s", errStatus.Code().String())

		go func() {
			ended := metric.Done(true)
			l.Debugf("rendering %s took %d ms", info.FullMethod, ended.Milliseconds())
		}()

		return nil, err
	}

	go func() {
		ended := metric.Done(false)
		l.Debugf("rendering %s took %d ms", info.FullMethod, ended.Milliseconds())
	}()

	return i, err
}
