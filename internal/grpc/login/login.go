package login

import (
	"github.com/feditools/login/internal/db"
	"github.com/feditools/login/internal/grpc"
	"github.com/feditools/login/pkg/proto"
)

// Module is a grpc login server module.
type Module struct {
	proto.UnimplementedLoginServer

	db db.DB
}

// New creates a new grpc login server modules.
func New(d db.DB) (*Module, error) {
	return &Module{
		db: d,
	}, nil
}

// Name return the module name.
func (Module) Name() string {
	return "login"
}

// Register registers the service with the grpc server.
func (m *Module) Register(s *grpc.Server) error {
	proto.RegisterLoginServer(s.Server(), m)

	return nil
}
