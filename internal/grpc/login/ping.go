package login

import (
	"context"
	"github.com/feditools/login/pkg/proto"
)

// Ping returns a "pong" response.
func (Module) Ping(_ context.Context, _ *proto.PingRequest) (*proto.Pong, error) {
	return &proto.Pong{Response: "pong"}, nil
}
