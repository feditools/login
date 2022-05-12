package grpc

import (
	"context"
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
		return nil, err
	}

	cancel()
	return resp, nil
}
