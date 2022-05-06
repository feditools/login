package redis

import (
	"context"
	"github.com/feditools/login/internal/kv"
	"time"
)

// fedi actor

// DeleteFediActor deletes fedi actor from redis.
func (c *Client) DeleteFediActor(ctx context.Context, uri string) error {
	_, err := c.redis.Del(ctx, kv.KeyFediActor(uri)).Result()
	if err != nil {
		return err
	}

	return nil
}

// GetFediActor retrieves fedi actor from redis.
func (c *Client) GetFediActor(ctx context.Context, uri string) (string, error) {
	resp, err := c.redis.Get(ctx, kv.KeyFediActor(uri)).Result()
	if err != nil {
		return "", err
	}

	return resp, nil
}

// SetFediActor adds fedi actor to redis.
func (c *Client) SetFediActor(ctx context.Context, uri string, actor string, expire time.Duration) error {
	_, err := c.redis.SetEX(ctx, kv.KeyFediActor(uri), actor, expire).Result()
	if err != nil {
		return err
	}

	return nil
}

// fedi node info

// DeleteFediNodeInfo deletes fedi nodeinfo from redis.
func (c *Client) DeleteFediNodeInfo(ctx context.Context, domain string) error {
	_, err := c.redis.Del(ctx, kv.KeyFediNodeInfo(domain)).Result()
	if err != nil {
		return err
	}

	return nil
}

// GetFediNodeInfo retrieves fedi nodeinfo from redis.
func (c *Client) GetFediNodeInfo(ctx context.Context, domain string) (string, error) {
	resp, err := c.redis.Get(ctx, kv.KeyFediNodeInfo(domain)).Result()
	if err != nil {
		return "", err
	}

	return resp, nil
}

// SetFediNodeInfo adds fedi nodeinfo to redis.
func (c *Client) SetFediNodeInfo(ctx context.Context, domain string, nodeinfo string, expire time.Duration) error {
	_, err := c.redis.SetEX(ctx, kv.KeyFediNodeInfo(domain), nodeinfo, expire).Result()
	if err != nil {
		return err
	}

	return nil
}
