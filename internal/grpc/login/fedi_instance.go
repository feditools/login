package login

import (
	"context"

	"github.com/feditools/login/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetFediInstance returns a federated social instance.
func (m Module) GetFediInstance(ctx context.Context, request *proto.GetFediInstanceRequest) (*proto.FediInstance, error) {
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

	return &proto.FediInstance{
		Id:             fediInstance.ID,
		Domain:         fediInstance.Domain,
		ServerHostname: fediInstance.ServerHostname,
		Software:       fediInstance.Software,
	}, nil
}
