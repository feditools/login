package kv

import (
	"context"
	"time"
)

// KV represents a key value store.
type KV interface {
	Close(ctx context.Context) error

	// federated actor

	DeleteFediActor(ctx context.Context, uri string) (err error)
	GetFediActor(ctx context.Context, uri string) (actor string, err error)
	SetFediActor(ctx context.Context, uri string, actor string, expire time.Duration) (err error)

	// federated instance node info

	DeleteFediNodeInfo(ctx context.Context, domain string) (err error)
	GetFediNodeInfo(ctx context.Context, domain string) (nodeinfo string, err error)
	SetFediNodeInfo(ctx context.Context, domain string, nodeinfo string, expire time.Duration) (err error)

	// oauth nonce

	DeleteOauthNonce(ctx context.Context, uid string, sid string) (err error)
	GetOauthNonce(ctx context.Context, uid string, sid string) (nonce string, err error)
	SetOauthNonce(ctx context.Context, uid string, sid string, nonce string, expire time.Duration) (err error)
}
