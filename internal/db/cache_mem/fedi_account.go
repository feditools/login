package cachemem

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/gob"
	"github.com/allegro/bigcache/v3"
	"github.com/feditools/login/internal/db"
	"github.com/feditools/login/internal/models"
	"strconv"
	"strings"
)

// CountFediAccounts returns the number of federated social account
func (c *CacheMem) CountFediAccounts(ctx context.Context) (int64, db.Error) {
	metric := c.metrics.NewDBCacheQuery("CountFediAccounts")

	count, hit := c.getCount(ctx, keyCountFediAccounts())
	if hit {
		go metric.Done(true, false)
		return count, nil
	}
	count, err := c.db.CountFediAccounts(ctx)
	if err != nil {
		go metric.Done(false, true)
		return 0, err
	}
	if count != 0 {
		c.setCount(ctx, keyCountFediAccounts(), count)
	}
	go metric.Done(false, false)
	return count, nil
}

// CountFediAccountsForInstance returns the number of federated social account for an instance
func (c *CacheMem) CountFediAccountsForInstance(ctx context.Context, instanceID int64) (int64, db.Error) {
	metric := c.metrics.NewDBCacheQuery("CountFediAccountsForInstance")

	count, hit := c.getCount(ctx, keyCountFediAccountsForInstance(instanceID))
	if hit {
		go metric.Done(true, false)
		return count, nil
	}
	count, err := c.db.CountFediAccountsForInstance(ctx, instanceID)
	if err != nil {
		go metric.Done(false, true)
		return 0, err
	}
	if count != 0 {
		c.setCount(ctx, keyCountFediAccountsForInstance(instanceID), count)
	}
	go metric.Done(false, false)
	return count, nil
}

// CreateFediAccount stores the federated instance and caches it
func (c *CacheMem) CreateFediAccount(ctx context.Context, account *models.FediAccount) db.Error {
	err := c.db.CreateFediAccount(ctx, account)
	if err != nil {
		return err
	}
	c.setFediAccount(ctx, account)
	return nil
}

// ReadFediAccount returns one federated social account
func (c *CacheMem) ReadFediAccount(ctx context.Context, id int64) (*models.FediAccount, db.Error) {
	metric := c.metrics.NewDBCacheQuery("ReadFediAccount")

	account, hit := c.getFediAccount(ctx, id)
	if hit {
		go metric.Done(true, false)
		return account, nil
	}
	account, err := c.db.ReadFediAccount(ctx, id)
	if err != nil {
		go metric.Done(false, true)
		return nil, err
	}
	if account != nil {
		c.setFediAccount(ctx, account)
	}
	go metric.Done(false, false)
	return account, nil
}

// ReadFediAccountByUsername returns one federated social account
func (c *CacheMem) ReadFediAccountByUsername(ctx context.Context, instanceID int64, username string) (*models.FediAccount, db.Error) {
	metric := c.metrics.NewDBCacheQuery("ReadFediAccountByUsername")

	account, hit := c.getFediAccountByUsername(ctx, instanceID, username)
	if hit {
		go metric.Done(true, false)
		return account, nil
	}
	account, err := c.db.ReadFediAccountByUsername(ctx, instanceID, username)
	if err != nil {
		go metric.Done(false, true)
		return nil, err
	}
	if account != nil {
		c.setFediAccount(ctx, account)
	}
	go metric.Done(false, false)
	return account, nil
}

// UpdateFediAccount updates the stored federated instance and caches it
func (c *CacheMem) UpdateFediAccount(ctx context.Context, account *models.FediAccount) db.Error {
	err := c.db.UpdateFediAccount(ctx, account)
	if err != nil {
		return err
	}
	c.setFediAccount(ctx, account)
	return nil
}

func (c *CacheMem) getFediAccount(_ context.Context, id int64) (*models.FediAccount, bool) {
	l := logger.WithField("func", "getFediAccount")

	// check cache
	entry, err := c.fediAccount.Get(strconv.FormatInt(id, 10))
	if err == bigcache.ErrEntryNotFound {
		return nil, false
	}
	if err != nil {
		l.Warnf("cache get: %s", err.Error())
		return nil, false
	}

	// decode
	buf := bytes.NewBuffer(entry)
	dec := gob.NewDecoder(buf)
	account := new(models.FediAccount)
	if err := dec.Decode(&account); err != nil {
		l.Warnf("cache decode: %s", err.Error())
		return nil, false
	}
	return account, true
}

func (c *CacheMem) getFediAccountByUsername(ctx context.Context, instanceID int64, username string) (*models.FediAccount, bool) {
	l := logger.WithField("func", "getFediAccountByUsername")

	// check username cache
	entry, err := c.fediAccountUsernameToID.Get(keyFediAccountByUsername(instanceID, username))
	if err == bigcache.ErrEntryNotFound {
		return nil, false
	}
	if err != nil {
		l.Warnf("cache get: %s", err.Error())
		return nil, false
	}
	i := int64(binary.LittleEndian.Uint64(entry))

	return c.getFediAccount(ctx, i)
}

func (c *CacheMem) setFediAccount(_ context.Context, account *models.FediAccount) {
	l := logger.WithField("func", "cacheSetFediAccount")

	// encode object
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(account); err != nil {
		l.Warnf("cache encode: %s", err.Error())
		return
	}

	// put in the cache
	err := c.fediAccount.Set(strconv.FormatInt(account.ID, 10), buf.Bytes())
	if err != nil {
		l.Warnf("cache obj: %s", err.Error())
		return
	}

	// encode id
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(account.ID))
	err = c.fediAccountUsernameToID.Set(keyFediAccountByUsername(account.InstanceID, account.Username), b)
	if err != nil {
		l.Warnf("cache domain: %s", err.Error())
		return
	}
}

func keyFediAccountByUsername(instanceID int64, username string) string {
	return strconv.FormatInt(instanceID, 10) + "|" + strings.ToLower(username)
}
