package grpc

import (
	"context"
)

// Ping sends a request to the server and response with "pong"
func (c *Client) Ping(ctx context.Context) (*PingReply, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)

	resp, err := c.ping.Ping(ctx, &PingRequest{})
	if err != nil {
		cancel()
		return nil, err
	}

	cancel()
	return resp, nil
}
