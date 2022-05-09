package ping

import (
	"context"
	"github.com/feditools/login/internal/grpc"
	pubgrpc "github.com/feditools/login/pkg/grpc"
)

type Ping interface {
	grpc.Module
	pubgrpc.PingServer
}

type Module struct {
	pubgrpc.UnimplementedPingServer
}

func New() (Ping, error) {
	return &Module{}, nil
}

// Ping returns a "pong" response
func (Module) Ping(_ context.Context, _ *pubgrpc.PingRequest) (*pubgrpc.PingReply, error) {
	return &pubgrpc.PingReply{Message: "pong"}, nil
}

// Name return the module name.
func (Module) Name() string {
	return "ping"
}

// Register registers the service with the grpc server.
func (m *Module) Register(s *grpc.Server) error {
	pubgrpc.RegisterPingServer(s.Server(), m)
	return nil
}
