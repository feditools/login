package ping

import (
	"context"
	"github.com/feditools/login/internal/grpc"
	pkg "github.com/feditools/login/pkg/grpc"
)

type Ping interface {
	grpc.Module
	pkg.PingServer
}

type Module struct {
	pkg.UnimplementedPingServer
}

func New() (Ping, error) {
	return &Module{}, nil
}

// Ping returns a "pong" response
func (Module) Ping(_ context.Context, _ *pkg.PingRequest) (*pkg.PingReply, error) {
	return &pkg.PingReply{Message: "pong"}, nil
}

// Name return the module name.
func (Module) Name() string {
	return "ping"
}

// Register registers the service with the grpc server.
func (m *Module) Register(s *grpc.Server) error {
	pkg.RegisterPingServer(s.Server(), m)
	return nil
}
