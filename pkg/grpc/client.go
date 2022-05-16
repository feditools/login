package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const defaultTimeout = 10 * time.Second

// Client is a feditools login grpc client
type Client struct {
	conn *grpc.ClientConn

	fediAccount  FediAccountClient
	fediInstance FediInstanceClient
	ping         PingClient
}

// NewClient creates a new feditools login grpc client
func NewClient(address string, cred credentials.PerRPCCredentials) (*Client, error) {
	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(cred),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// services
	fediAccountC := NewFediAccountClient(conn)
	fediInstanceC := NewFediInstanceClient(conn)
	pingC := NewPingClient(conn)

	return &Client{
		conn: conn,

		fediAccount:  fediAccountC,
		fediInstance: fediInstanceC,
		ping:         pingC,
	}, nil
}

// Close closes the feditools login grpc client
func (c *Client) Close() error {
	return c.conn.Close()
}
