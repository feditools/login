package grpc

import (
	"context"

	"github.com/feditools/login/pkg/proto"
)

// Ping sends a request to the server and response with "pong".
func (c *Client) Ping(ctx context.Context) (*proto.Pong, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)

	resp, err := c.login.Ping(ctx, &proto.PingRequest{})
	if err != nil {
		cancel()

		return nil, err
	}

	cancel()

	return resp, nil
}
