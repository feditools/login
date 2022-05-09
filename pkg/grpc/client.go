package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

// Client is a feditools login grpc client
type Client struct {
	conn *grpc.ClientConn

	ping PingClient
}

// NewClient creates a new feditools login grpc client
func NewClient(address string, cred credentials.PerRPCCredentials) (*Client, error) {
	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(cred),
	}

	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// services
	pingC := NewPingClient(conn)

	return &Client{
		conn: conn,

		ping: pingC,
	}, nil
}

// Close closes the feditools login grpc client
func (c *Client) Close() error {
	return c.conn.Close()
}
