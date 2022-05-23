package grpc

import (
	"time"

	"github.com/feditools/login/pkg/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

const defaultTimeout = 10 * time.Second

// Client is a feditools login grpc client.
type Client struct {
	conn  *grpc.ClientConn
	login proto.LoginClient
}

// NewClient creates a new feditools login grpc client.
func NewClient(address string, cred credentials.PerRPCCredentials) (*Client, error) {
	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(cred),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		return nil, err
	}

	// services
	loginC := proto.NewLoginClient(conn)

	return &Client{
		conn:  conn,
		login: loginC,
	}, nil
}

// Close closes the feditools login grpc client.
func (c *Client) Close() error {
	return c.conn.Close()
}
