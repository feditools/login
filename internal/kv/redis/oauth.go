package redis

import (
	"context"
	"github.com/feditools/login/internal/kv"
	"time"
)

// oauth nonce

// DeleteOauthNonce deletes an oauth nonce from redis.
func (c *Client) DeleteOauthNonce(ctx context.Context, uid int64, sid string) error {
	_, err := c.redis.Del(ctx, kv.KeyOauthNonce(uid, sid)).Result()
	if err != nil {
		return err
	}

	return nil
}

// GetOauthNonce retrieves an oauth nonce from redis.
func (c *Client) GetOauthNonce(ctx context.Context, uid int64, sid string) (string, error) {
	resp, err := c.redis.Get(ctx, kv.KeyOauthNonce(uid, sid)).Result()
	if err != nil {
		return "", err
	}

	return resp, nil
}

// SetOauthNonce adds an oauth nonce to redis.
func (c *Client) SetOauthNonce(ctx context.Context, uid int64, sid string, nonce string, expire time.Duration) error {
	_, err := c.redis.SetEX(ctx, kv.KeyOauthNonce(uid, sid), nonce, expire).Result()
	if err != nil {
		return err
	}

	return nil
}
