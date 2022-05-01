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

// CreateFediInstance stores the federated instance and caches it
func (c *CacheMem) CreateFediInstance(ctx context.Context, instance *models.FediInstance) db.Error {
	err := c.db.CreateFediInstance(ctx, instance)
	if err != nil {
		return err
	}
	c.setFediInstance(ctx, instance)
	return nil
}

// ReadFediInstance returns one federated social instance
func (c *CacheMem) ReadFediInstance(ctx context.Context, id int64) (*models.FediInstance, db.Error) {
	metric := c.metrics.NewDBCacheQuery("ReadFediInstance")

	instance, hit := c.getFediInstance(ctx, id)
	if hit {
		go metric.Done(true, false)
		return instance, nil
	}
	instance, err := c.db.ReadFediInstance(ctx, id)
	if err != nil {
		go metric.Done(false, true)
		return nil, err
	}
	c.setFediInstance(ctx, instance)
	go metric.Done(false, false)
	return instance, nil
}

// ReadFediInstanceByDomain returns one federated social instance
func (c *CacheMem) ReadFediInstanceByDomain(ctx context.Context, domain string) (*models.FediInstance, db.Error) {
	metric := c.metrics.NewDBCacheQuery("ReadFediInstanceByDomain")

	instance, hit := c.getFediInstanceByDomain(ctx, domain)
	if hit {
		go metric.Done(true, false)
		return instance, nil
	}
	instance, err := c.db.ReadFediInstanceByDomain(ctx, domain)
	if err != nil {
		go metric.Done(false, true)
		return nil, err
	}
	c.setFediInstance(ctx, instance)
	go metric.Done(false, false)
	return instance, nil
}

// ReadFediInstancesPage returns a page of federated social instances
func (c *CacheMem) ReadFediInstancesPage(ctx context.Context, index, count int) ([]*models.FediInstance, db.Error) {
	return c.db.ReadFediInstancesPage(ctx, index, count)
}

// UpdateFediInstance updates the stored federated instance and caches it
func (c *CacheMem) UpdateFediInstance(ctx context.Context, instance *models.FediInstance) db.Error {
	err := c.db.CreateFediInstance(ctx, instance)
	if err != nil {
		return err
	}
	c.setFediInstance(ctx, instance)
	return nil
}

func (c *CacheMem) getFediInstance(_ context.Context, id int64) (*models.FediInstance, bool) {
	l := logger.WithField("func", "getFediInstance")

	// check cache
	entry, err := c.fediInstance.Get(strconv.FormatInt(id, 10))
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
	instance := new(models.FediInstance)
	if err := dec.Decode(&instance); err != nil {
		l.Warnf("cache decode: %s", err.Error())
		return nil, false
	}
	return instance, true
}

func (c *CacheMem) getFediInstanceByDomain(ctx context.Context, domain string) (*models.FediInstance, bool) {
	l := logger.WithField("func", "getFediInstanceByDomain")

	// check domain cache
	entry, err := c.fediInstanceDomainToID.Get(strings.ToLower(domain))
	if err == bigcache.ErrEntryNotFound {
		return nil, false
	}
	if err != nil {
		l.Warnf("cache get: %s", err.Error())
		return nil, false
	}
	i := int64(binary.LittleEndian.Uint64(entry))

	return c.getFediInstance(ctx, i)
}

func (c *CacheMem) setFediInstance(_ context.Context, instance *models.FediInstance) {
	l := logger.WithField("func", "setFediInstance")

	// encode object
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(instance); err != nil {
		l.Warnf("cache encode: %s", err.Error())
		return
	}

	// put in the cache
	err := c.fediInstance.Set(strconv.FormatInt(instance.ID, 10), buf.Bytes())
	if err != nil {
		l.Warnf("cache obj: %s", err.Error())
		return
	}

	// encode id
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(instance.ID))
	err = c.fediInstanceDomainToID.Set(strings.ToLower(instance.Domain), b)
	if err != nil {
		l.Warnf("cache domain: %s", err.Error())
		return
	}
}
