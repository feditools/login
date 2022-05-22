package redis

import (
	"context"
	"time"

	"github.com/feditools/login/internal/kv"
)

// oauth nonce

// DeleteOauthNonce deletes an oauth nonce from redis.
func (c *Client) DeleteOauthNonce(ctx context.Context, uid string, sid string) error {
	l := logger.WithField("func", "DeleteOauthNonce")
	l.Tracef("DeleteOauthNonce(ctx, %s, %s) called", uid, sid)

	_, err := c.redis.Del(ctx, kv.KeyOauthNonce(uid, sid)).Result()
	if err != nil {
		l.Tracef("DeleteOauthNonce(ctx, %s, %s) returned (%s %T)", uid, sid, err.Error(), err)

		return err
	}

	l.Tracef("DeleteOauthNonce(ctx, %s, %s) returned (nil nil)", uid, sid)

	return nil
}

// GetOauthNonce retrieves an oauth nonce from redis.
func (c *Client) GetOauthNonce(ctx context.Context, uid string, sid string) (string, error) {
	l := logger.WithField("func", "GetOauthNonce")
	l.Tracef("GetOauthNonce(ctx, %s, %s) called", uid, sid)

	resp, err := c.redis.Get(ctx, kv.KeyOauthNonce(uid, sid)).Result()
	if err != nil {
		l.Tracef("GetOauthNonce(ctx, %s, %s) returned (\"\" string, %s %T)", uid, sid, err.Error(), err)

		return "", err
	}

	l.Tracef("GetOauthNonce(ctx, %s, %s) returned (\"%s\" string, nil nil)", uid, sid, resp)

	return resp, nil
}

// SetOauthNonce adds an oauth nonce to redis.
func (c *Client) SetOauthNonce(ctx context.Context, uid string, sid string, nonce string, expire time.Duration) error {
	l := logger.WithField("func", "SetOauthNonce")
	l.Tracef("SetOauthNonce(ctx, %s, %s, %s, %s) called", uid, sid, nonce, expire.String())

	_, err := c.redis.SetEX(ctx, kv.KeyOauthNonce(uid, sid), nonce, expire).Result()
	if err != nil {
		l.Tracef("SetOauthNonce(ctx, %s, %s, %s, %s) returned (%s %T)", uid, sid, nonce, expire.String(), err.Error(), err)

		return err
	}

	l.Tracef("SetOauthNonce(ctx, %s, %s, %s, %s) returned (nil nil)", uid, sid, nonce, expire.String())

	return nil
}
