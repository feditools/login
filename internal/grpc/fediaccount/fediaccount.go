package fediaccount

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
	pkg.UnimplementedFediAccountServer

	db db.DB
}

// New creates a new grpc ping server modules
func New(d db.DB) (*Module, error) {
	return &Module{
		db: d,
	}, nil
}

// GetFediAccount returns a federated account
func (m Module) GetFediAccount(ctx context.Context, request *pkg.GetFediAccountRequest) (*pkg.GetFediAccountReply, error) {
	l := logger.WithField("func", "GetFediAccount")

	fediAccount, err := m.db.ReadFediAccount(ctx, request.Id)
	if err != nil {
		l.Errorf("db read: %s", err.Error())
		return nil, err
	}
	if fediAccount == nil {
		st := status.New(codes.NotFound, "FediAccount not found.")
		return nil, st.Err()
	}

	return &pkg.GetFediAccountReply{
		Id:          fediAccount.ID,
		Username:    fediAccount.Username,
		InstanceId:  fediAccount.InstanceID,
		DisplayName: fediAccount.DisplayName,
		IsAdmin:     fediAccount.Admin,
	}, nil
}

// Name return the module name.
func (Module) Name() string {
	return "ping"
}

// Register registers the service with the grpc server.
func (m *Module) Register(s *grpc.Server) error {
	pkg.RegisterFediAccountServer(s.Server(), m)
	return nil
}
