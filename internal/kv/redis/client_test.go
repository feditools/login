//go:build redis

package redis

import (
	"context"
	"github.com/feditools/login/internal/config"
	"github.com/spf13/viper"
	"testing"
)

func TestNew(t *testing.T) {
	viper.Reset()

	viper.Set(config.Keys.RedisAddress, "redis:6379")
	viper.Set(config.Keys.RedisPassword, "test")

	client, err := New(context.Background())
	if err != nil {
		t.Errorf("unexpected error initializing redis connection: %s", err.Error())
		return
	}
	if client == nil {
		t.Errorf("client is nil")
		return
	}
}

func TestClient_Close(t *testing.T) {
	viper.Reset()

	viper.Set(config.Keys.RedisAddress, "redis:6379")
	viper.Set(config.Keys.RedisPassword, "test")

	client, err := New(context.Background())
	if err != nil {
		t.Errorf("unexpected error initializing redis connection: %s", err.Error())
		return
	}

	err = client.Close(context.Background())
	if err != nil {
		t.Errorf("unexpected error closing redis connection: %s", err.Error())
		return
	}
}

func TestClient_RedisClient(t *testing.T) {
	viper.Reset()

	viper.Set(config.Keys.RedisAddress, "redis:6379")
	viper.Set(config.Keys.RedisPassword, "test")

	client, err := New(context.Background())
	if err != nil {
		t.Errorf("unexpected error initializing redis connection: %s", err.Error())
		return
	}

	r := client.RedisClient()
	if r == nil {
		t.Errorf("redis client is nil")
		return
	}

}
