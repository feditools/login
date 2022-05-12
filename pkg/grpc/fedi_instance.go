package grpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetFediInstance retrieves a federated account from the login server
func (c *Client) GetFediInstance(ctx context.Context, id int64) (*GetFediInstanceReply, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)

	req := &GetFediInstanceRequest{
		Id: id,
	}

	resp, err := c.fediInstance.GetFediInstance(ctx, req)
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
