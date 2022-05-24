package redis

import (
	"context"
	"time"

	"github.com/feditools/login/internal/kv"
)

// oauth nonce login

// DeleteOauthNonceLogin deletes an oauth nonce from redis.
func (c *Client) DeleteOauthNonceLogin(ctx context.Context, uid string) error {
	l := logger.WithField("func", "DeleteOauthNonceLogin")
	l.Tracef("DeleteOauthNonceLogin(ctx, %s) called", uid)

	_, err := c.redis.Del(ctx, kv.KeyOauthNonceLogin(uid)).Result()
	if err != nil {
		l.Tracef("DeleteOauthNonceLogin(ctx, %s) returned (%s %T)", uid, err.Error(), err)

		return err
	}

	l.Tracef("DeleteOauthNonceLogin(ctx, %s) returned (nil nil)", uid)

	return nil
}

// GetOauthNonceLogin retrieves an oauth nonce from redis.
func (c *Client) GetOauthNonceLogin(ctx context.Context, uid string) (string, error) {
	l := logger.WithField("func", "GetOauthNonceLogin")
	l.Tracef("GetOauthNonceLogin(ctx, %s) called", uid)

	resp, err := c.redis.Get(ctx, kv.KeyOauthNonceLogin(uid)).Result()
	if err != nil {
		l.Tracef("GetOauthNonceLogin(ctx, %s) returned (\"\" string, %s %T)", uid, err.Error(), err)

		return "", err
	}

	l.Tracef("GetOauthNonceLogin(ctx, %s) returned (\"%s\" string, nil nil)", uid, resp)

	return resp, nil
}

// SetOauthNonceLogin adds an oauth nonce to redis.
func (c *Client) SetOauthNonceLogin(ctx context.Context, uid string, nonce string, expire time.Duration) error {
	l := logger.WithField("func", "SetOauthNonceLogin")
	l.Tracef("SetOauthNonceLogin(ctx, %s, %s, %s) called", uid, nonce, expire.String())

	_, err := c.redis.SetEX(ctx, kv.KeyOauthNonceLogin(uid), nonce, expire).Result()
	if err != nil {
		l.Tracef("SetOauthNonceLogin(ctx, %s, %s, %s) returned (%s %T)", uid, nonce, expire.String(), err.Error(), err)

		return err
	}

	l.Tracef("SetOauthNonceLogin(ctx, %s, %s, %s) returned (nil nil)", uid, nonce, expire.String())

	return nil
}

// oauth nonce refresh

// DeleteOauthNonceRefresh deletes an oauth nonce from redis.
func (c *Client) DeleteOauthNonceRefresh(ctx context.Context, refreshToken string) error {
	l := logger.WithField("func", "DeleteOauthNonceRefresh")
	l.Tracef("DeleteOauthNonceRefresh(ctx, %s) called", refreshToken)

	_, err := c.redis.Del(ctx, kv.KeyOauthNonceRefresh(refreshToken)).Result()
	if err != nil {
		l.Tracef("DeleteOauthNonceRefresh(ctx, %s) returned (%s %T)", refreshToken, err.Error(), err)

		return err
	}

	l.Tracef("DeleteOauthNonceRefresh(ctx, %s) returned (nil nil)", refreshToken)

	return nil
}

// GetOauthNonceRefresh retrieves an oauth nonce from redis.
func (c *Client) GetOauthNonceRefresh(ctx context.Context, refreshToken string) (string, error) {
	l := logger.WithField("func", "GetOauthNonceRefresh")
	l.Tracef("GetOauthNonceRefresh(ctx, %s) called", refreshToken)

	resp, err := c.redis.Get(ctx, kv.KeyOauthNonceRefresh(refreshToken)).Result()
	if err != nil {
		l.Tracef("GetOauthNonceRefresh(ctx, %s) returned (\"\" string, %s %T)", refreshToken, err.Error(), err)

		return "", err
	}

	l.Tracef("GetOauthNonceRefresh(ctx, %s) returned (\"%s\" string, nil nil)", refreshToken, resp)

	return resp, nil
}

// SetOauthNonceRefresh adds an oauth nonce to redis.
func (c *Client) SetOauthNonceRefresh(ctx context.Context, refreshToken string, nonce string, expire time.Duration) error {
	l := logger.WithField("func", "SetOauthNonceRefresh")
	l.Tracef("SetOauthNonceRefresh(ctx, %s, %s, %s) called", refreshToken, nonce, expire.String())

	_, err := c.redis.SetEX(ctx, kv.KeyOauthNonceRefresh(refreshToken), nonce, expire).Result()
	if err != nil {
		l.Tracef("SetOauthNonceRefresh(ctx, %s, %s, %s) returned (%s %T)", refreshToken, nonce, expire.String(), err.Error(), err)

		return err
	}

	l.Tracef("SetOauthNonceRefresh(ctx, %s, %s, %s) returned (nil nil)", refreshToken, nonce, expire.String())

	return nil
}
