package grpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
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
			ended := metric.Done(int(codes.Internal))
			l.Debugf("rendering %s took %d ms", info.FullMethod, ended.Milliseconds())
		}()

		st := status.New(codes.Internal, "Missing metadata.")
		return nil, st.Err()
	}

	// validate login
	if !s.authValid(ctx, md["authorization"]) {
		go func() {
			ended := metric.Done(int(codes.Unauthenticated))
			l.Debugf("rendering %s took %d ms", info.FullMethod, ended.Milliseconds())
		}()

		st := status.New(codes.Unauthenticated, "Invalid token.")
		return nil, st.Err()
	}

	// do request
	i, err := handler(ctx, req)
	if err != nil {
		respStatuc := status.Convert(err)

		go func() {
			ended := metric.Done(int(respStatuc.Code()))
			l.Debugf("rendering %s took %d ms", info.FullMethod, ended.Milliseconds())
		}()

		return nil, err
	}

	go func() {
		ended := metric.Done(int(codes.OK))
		l.Debugf("rendering %s took %d ms", info.FullMethod, ended.Milliseconds())
	}()

	return i, err
}
