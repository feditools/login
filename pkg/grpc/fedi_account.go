package grpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetFediAccount retrieves a federated account from the login server
func (c *Client) GetFediAccount(ctx context.Context, id int64) (*GetFediAccountReply, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)

	req := &GetFediAccountRequest{
		Id: id,
	}

	resp, err := c.fediAccount.GetFediAccount(ctx, req)
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
