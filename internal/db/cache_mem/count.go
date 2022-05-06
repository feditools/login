package cachemem

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/allegro/bigcache/v3"
)

const (
	tableNameFediAccounts  = "fedi_accounts"
	tableNameFediInstances = "fedi_instances"
	tableNameOauthClients  = "oauth_clients"
)

func keyCountOauthClients() string {
	return tableNameOauthClients
}
func keyCountFediAccounts() string {
	return tableNameFediAccounts
}
func keyCountFediAccountsForInstance(instanceID int64) string {
	return fmt.Sprintf("%s-i-%d", tableNameFediAccounts, instanceID)
}
func keyCountFediInstances() string {
	return tableNameFediInstances
}

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
