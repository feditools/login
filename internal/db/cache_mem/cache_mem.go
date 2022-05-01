package cachemem

import (
	"context"
	bigcache "github.com/allegro/bigcache/v3"
	"github.com/feditools/login/internal/db"
	"github.com/feditools/login/internal/metrics"
	"time"
)

// CacheMem is an in memory caching middleware for our db interface
type CacheMem struct {
	db      db.DB
	metrics metrics.Collector

	fediAccount             *bigcache.BigCache
	fediAccountUsernameToID *bigcache.BigCache

	fediInstance           *bigcache.BigCache
	fediInstanceDomainToID *bigcache.BigCache
}

// New creates a new in memory cache
func New(_ context.Context, d db.DB, m metrics.Collector) (db.DB, error) {
	fediAccount, err := bigcache.NewBigCache(bigcache.Config{
		Shards:             32,
		LifeWindow:         10 * time.Minute,
		CleanWindow:        5 * time.Minute,
		MaxEntriesInWindow: 10000,
		MaxEntrySize:       500,
		Verbose:            true,
		HardMaxCacheSize:   8192,
	})
	if err != nil {
		return nil, err
	}

	fediAccountUsernameToID, err := bigcache.NewBigCache(bigcache.Config{
		Shards:             32,
		LifeWindow:         10 * time.Minute,
		CleanWindow:        5 * time.Minute,
		MaxEntriesInWindow: 10000,
		MaxEntrySize:       8,
		Verbose:            true,
		HardMaxCacheSize:   8192,
	})
	if err != nil {
		return nil, err
	}

	fediInstance, err := bigcache.NewBigCache(bigcache.Config{
		Shards:             32,
		LifeWindow:         10 * time.Minute,
		CleanWindow:        5 * time.Minute,
		MaxEntriesInWindow: 10000,
		MaxEntrySize:       500,
		Verbose:            true,
		HardMaxCacheSize:   8192,
	})
	if err != nil {
		return nil, err
	}

	fediInstanceDomainToID, err := bigcache.NewBigCache(bigcache.Config{
		Shards:             32,
		LifeWindow:         10 * time.Minute,
		CleanWindow:        5 * time.Minute,
		MaxEntriesInWindow: 10000,
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

		fediAccount:             fediAccount,
		fediAccountUsernameToID: fediAccountUsernameToID,

		fediInstance:           fediInstance,
		fediInstanceDomainToID: fediInstanceDomainToID,
	}, nil
}

// Close is a pass through
func (c *CacheMem) Close(ctx context.Context) db.Error {
	return c.db.Close(ctx)
}

// Create is a pass through
func (c *CacheMem) Create(ctx context.Context, i any) db.Error {
	return c.db.Create(ctx, i)
}

// DoMigration is a pass through
func (c *CacheMem) DoMigration(ctx context.Context) db.Error {
	return c.db.DoMigration(ctx)
}

// LoadTestData is a pass through
func (c *CacheMem) LoadTestData(ctx context.Context) db.Error {
	return c.db.LoadTestData(ctx)
}

// ReadByID is a pass through
func (c *CacheMem) ReadByID(ctx context.Context, id int64, i any) db.Error {
	return c.db.ReadByID(ctx, id, i)
}

// Update is a pass through
func (c *CacheMem) Update(ctx context.Context, i any) db.Error {
	return c.db.Update(ctx, i)
}
