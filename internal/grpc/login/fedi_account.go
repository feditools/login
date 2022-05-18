package login

import (
	"context"
	"github.com/feditools/login/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetFediAccount returns a federated account.
func (m Module) GetFediAccount(ctx context.Context, request *proto.GetFediAccountRequest) (*proto.FediAccount, error) {
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

	return &proto.FediAccount{
		Id:          fediAccount.ID,
		Username:    fediAccount.Username,
		InstanceId:  fediAccount.InstanceID,
		DisplayName: fediAccount.DisplayName,
		IsAdmin:     fediAccount.Admin,
	}, nil
}
