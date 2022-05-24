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

	// oauth nonce login

	DeleteOauthNonceLogin(ctx context.Context, uid string) (err error)
	GetOauthNonceLogin(ctx context.Context, uid string) (nonce string, err error)
	SetOauthNonceLogin(ctx context.Context, uid string, nonce string, expire time.Duration) (err error)

	// oauth nonce refresh

	DeleteOauthNonceRefresh(ctx context.Context, refreshToken string) (err error)
	GetOauthNonceRefresh(ctx context.Context, refreshToken string) (nonce string, err error)
	SetOauthNonceRefresh(ctx context.Context, refreshToken string, nonce string, expire time.Duration) (err error)
}
