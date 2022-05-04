package cachemem

import (
	"context"
	"encoding/binary"
	"github.com/allegro/bigcache/v3"
)

const (
	tableNameOauthClients = "oauth_clients"
)

func keyCountFullTable(s string) string { return s }

func (c *CacheMem) getCount(_ context.Context, k string) (int64, bool) {
	l := logger.WithField("func", "getCount")

	// check domain cache
	entry, err := c.fediInstanceDomainToID.Get(k)
	if err == bigcache.ErrEntryNotFound {
		return 0, false
	}
	if err != nil {
		l.Warnf("cache get: %s", err.Error())
		return 0, false
	}
	i := int64(binary.LittleEndian.Uint64(entry))

	return i, true
}

func (c *CacheMem) setCount(_ context.Context, k string, count int64) {
	l := logger.WithField("func", "setCount")

	// encode count
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(count))
	err := c.fediInstanceDomainToID.Set(k, b)
	if err != nil {
		l.Warnf("cache domain: %s", err.Error())
		return
	}
}
