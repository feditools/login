package fediinstance

import (
	"context"
	"github.com/feditools/login/internal/db"
	"github.com/feditools/login/internal/grpc"
	pkg "github.com/feditools/login/pkg/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Module is a grpc ping server module
type Module struct {
	pkg.UnimplementedFediInstanceServer

	db db.DB
}

// New creates a new grpc ping server modules
func New(d db.DB) (*Module, error) {
	return &Module{
		db: d,
	}, nil
}

// GetFediInstance returns a federated social instance
func (m Module) GetFediInstance(ctx context.Context, request *pkg.GetFediInstanceRequest) (*pkg.GetFediInstanceReply, error) {
	l := logger.WithField("func", "GetFediInstance")

	fediInstance, err := m.db.ReadFediInstance(ctx, request.Id)
	if err != nil {
		l.Errorf("db read: %s", err.Error())
		return nil, err
	}
	if fediInstance == nil {
		st := status.New(codes.NotFound, "FediInstance not found.")
		return nil, st.Err()
	}

	return &pkg.GetFediInstanceReply{
		Id:             fediInstance.ID,
		Domain:         fediInstance.Domain,
		ServerHostname: fediInstance.ServerHostname,
		Software:       fediInstance.Software,
	}, nil
}

// Name return the module name.
func (Module) Name() string {
	return "fedi_instance"
}

// Register registers the service with the grpc server.
func (m *Module) Register(s *grpc.Server) error {
	pkg.RegisterFediInstanceServer(s.Server(), m)
	return nil
}
