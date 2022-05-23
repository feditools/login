package grpc

import (
	"context"

	"github.com/feditools/login/pkg/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetFediInstance retrieves a federated account from the login server.
func (c *Client) GetFediInstance(ctx context.Context, id int64) (*proto.FediInstance, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)

	req := &proto.GetFediInstanceRequest{
		Id: id,
	}

	resp, err := c.login.GetFediInstance(ctx, req)
	if err != nil {
		cancel()
		respStatus := status.Convert(err)
		if respStatus.Code() == codes.NotFound {
			return nil, nil
		}

		return nil, err
	}

	cancel()

	return resp, nil
}
