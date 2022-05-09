package ping

import (
	"context"
	"github.com/feditools/login/internal/grpc"
)

type Ping interface {
	grpc.Module
	PingServer
}

type Module struct {
	UnimplementedPingServer
}

func New() (Ping, error) {
	return &Module{}, nil
}

// Ping returns a "pong" response
func (Module) Ping(_ context.Context, _ *PingRequest) (*PingReply, error) {
	return &PingReply{Message: "pong"}, nil
}

// Name return the module name.
func (Module) Name() string {
	return "ping"
}

// Register registers the service with the grpc server.
func (m *Module) Register(s *grpc.Server) error {
	RegisterPingServer(s.Server(), m)
	return nil
}
