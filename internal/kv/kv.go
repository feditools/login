package kv

import (
	"context"
	"time"
)

// KV represents a key value store
type KV interface {
	Close(ctx context.Context) error

	// federated instance node info

	DeleteFediNodeInfo(ctx context.Context, domain string) (err error)
	GetFediNodeInfo(ctx context.Context, domain string) (nodeinfo string, err error)
	SetFediNodeInfo(ctx context.Context, domain string, nodeinfo string, expire time.Duration) (err error)
}
