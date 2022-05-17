package cachemem

import (
	"context"
	"encoding/gob"
	"time"

	bigcache "github.com/allegro/bigcache/v3"
	"github.com/feditools/go-lib/metrics"
	"github.com/feditools/login/internal/db"
	"github.com/feditools/login/internal/models"
)

const (
	applicationTokenLifeWindow         = 10 * time.Minute
	applicationTokenCleanWindow        = 5 * time.Minute
	applicationTokenMaxEntriesInWindow = 1000

	fediAccountLifeWindow         = 10 * time.Minute
	fediAccountCleanWindow        = 5 * time.Minute
	fediAccountMaxEntriesInWindow = 10000

	fediInstanceLifeWindow         = 10 * time.Minute
	fediInstanceCleanWindow        = 5 * time.Minute
	fediInstanceMaxEntriesInWindow = 10000
)

// CacheMem is an in memory caching middleware for our db interface.
type CacheMem struct {
	db      db.DB
	metrics metrics.Collector

	count *bigcache.BigCache

	applicationToken          *bigcache.BigCache
	applicationTokenTokenToID *bigcache.BigCache

	fediAccount             *bigcache.BigCache
	fediAccountUsernameToID *bigcache.BigCache

	fediInstance           *bigcache.BigCache
	fediInstanceDomainToID *bigcache.BigCache

	allCaches []*bigcache.BigCache
}

// New creates a new in memory cache.
func New(_ context.Context, d db.DB, m metrics.Collector) (db.DB, error) {
	gob.Register(models.ApplicationToken{})
	gob.Register(models.FediAccount{})
	gob.Register(models.FediInstance{})

	count, err := bigcache.NewBigCache(bigcache.Config{
		Shards:             32,
		LifeWindow:         30 * time.Second,
		CleanWindow:        1 * time.Minute,
		MaxEntriesInWindow: 10000,
		MaxEntrySize:       8,
		Verbose:            true,
		HardMaxCacheSize:   8192,
	})
	if err != nil {
		return nil, err
	}

	applicationToken, err := bigcache.NewBigCache(bigcache.Config{
		Shards:             32,
		LifeWindow:         applicationTokenLifeWindow,
		CleanWindow:        applicationTokenCleanWindow,
		MaxEntriesInWindow: applicationTokenMaxEntriesInWindow,
		MaxEntrySize:       500,
		Verbose:            true,
		HardMaxCacheSize:   8192,
	})
	if err != nil {
		return nil, err
	}

	applicationTokenTokenToID, err := bigcache.NewBigCache(bigcache.Config{
		Shards:             32,
		LifeWindow:         applicationTokenLifeWindow,
		CleanWindow:        applicationTokenCleanWindow,
		MaxEntriesInWindow: applicationTokenMaxEntriesInWindow,
		MaxEntrySize:       8,
		Verbose:            true,
		HardMaxCacheSize:   8192,
	})
	if err != nil {
		return nil, err
	}

	fediAccount, err := bigcache.NewBigCache(bigcache.Config{
		Shards:             32,
		LifeWindow:         fediAccountLifeWindow,
		CleanWindow:        fediAccountCleanWindow,
		MaxEntriesInWindow: fediAccountMaxEntriesInWindow,
		MaxEntrySize:       500,
		Verbose:            true,
		HardMaxCacheSize:   8192,
	})
	if err != nil {
		return nil, err
	}

	fediAccountUsernameToID, err := bigcache.NewBigCache(bigcache.Config{
		Shards:             32,
		LifeWindow:         fediAccountLifeWindow,
		CleanWindow:        fediAccountCleanWindow,
		MaxEntriesInWindow: fediAccountMaxEntriesInWindow,
		MaxEntrySize:       8,
		Verbose:            true,
		HardMaxCacheSize:   8192,
	})
	if err != nil {
		return nil, err
	}

	fediInstance, err := bigcache.NewBigCache(bigcache.Config{
		Shards:             32,
		LifeWindow:         fediInstanceLifeWindow,
		CleanWindow:        fediInstanceCleanWindow,
		MaxEntriesInWindow: fediInstanceMaxEntriesInWindow,
		MaxEntrySize:       500,
		Verbose:            true,
		HardMaxCacheSize:   8192,
	})
	if err != nil {
		return nil, err
	}

	fediInstanceDomainToID, err := bigcache.NewBigCache(bigcache.Config{
		Shards:             32,
		LifeWindow:         fediInstanceLifeWindow,
		CleanWindow:        fediInstanceCleanWindow,
		MaxEntriesInWindow: fediInstanceMaxEntriesInWindow,
		MaxEntrySize:       8,
		Verbose:            true,
		HardMaxCacheSize:   8192,
	})
	if err != nil {
		return nil, err
	}

	return &CacheMem{
		db:      d,
		metrics: m,

		count: count,

		applicationToken:          applicationToken,
		applicationTokenTokenToID: applicationTokenTokenToID,

		fediAccount:             fediAccount,
		fediAccountUsernameToID: fediAccountUsernameToID,

		fediInstance:           fediInstance,
		fediInstanceDomainToID: fediInstanceDomainToID,

		allCaches: []*bigcache.BigCache{
			count,
			applicationToken,
			applicationTokenTokenToID,
			fediAccount,
			fediAccountUsernameToID,
			fediInstance,
			fediInstanceDomainToID,
		},
	}, nil
}

// Close is a pass through.
func (c *CacheMem) Close(ctx context.Context) db.Error {
	for _, cache := range c.allCaches {
		_ = cache.Close()
	}

	return c.db.Close(ctx)
}

// Create is a pass through.
func (c *CacheMem) Create(ctx context.Context, i interface{}) db.Error {
	return c.db.Create(ctx, i)
}

// DoMigration is a pass through.
func (c *CacheMem) DoMigration(ctx context.Context) db.Error {
	return c.db.DoMigration(ctx)
}

// LoadTestData is a pass through.
func (c *CacheMem) LoadTestData(ctx context.Context) db.Error {
	return c.db.LoadTestData(ctx)
}

// ReadByID is a pass through.
func (c *CacheMem) ReadByID(ctx context.Context, id int64, i interface{}) db.Error {
	return c.db.ReadByID(ctx, id, i)
}

// ResetCache clears all the caches.
func (c *CacheMem) ResetCache(ctx context.Context) db.Error {
	for _, cache := range c.allCaches {
		_ = cache.Reset()
	}

	return c.db.ResetCache(ctx)
}

// Update is a pass through.
func (c *CacheMem) Update(ctx context.Context, i interface{}) db.Error {
	return c.db.Update(ctx, i)
}
