package cachemem

import (
	"context"
	"github.com/feditools/login/internal/db"
	"github.com/feditools/login/internal/models"
)

// CountOauthClients returns the number of oauth clients
func (c *CacheMem) CountOauthClients(ctx context.Context) (int64, db.Error) {
	metric := c.metrics.NewDBCacheQuery("CountOauthClients")

	count, hit := c.getCount(ctx, keyCountOauthClients())
	if hit {
		go metric.Done(true, false)
		return count, nil
	}
	count, err := c.db.CountOauthClients(ctx)
	if err != nil {
		go metric.Done(false, true)
		return 0, err
	}
	if count != 0 {
		c.setCount(ctx, keyCountOauthClients(), count)
	}
	go metric.Done(false, false)
	return count, nil
}

// CreateOauthClient stores the oauth client
func (c *CacheMem) CreateOauthClient(ctx context.Context, client *models.OauthClient) (err db.Error) {
	return c.db.CreateOauthClient(ctx, client)
}

// ReadOauthClient returns one oauth client
func (c *CacheMem) ReadOauthClient(ctx context.Context, id int64) (*models.OauthClient, db.Error) {
	return c.db.ReadOauthClient(ctx, id)
}

// ReadOauthClientsPage returns a page of oauth clients
func (c *CacheMem) ReadOauthClientsPage(ctx context.Context, index, count int) (clients []*models.OauthClient, err db.Error) {
	return c.db.ReadOauthClientsPage(ctx, index, count)
}

// UpdateOauthClient updates the stored oauth client
func (c *CacheMem) UpdateOauthClient(ctx context.Context, client *models.OauthClient) (err db.Error) {
	return c.db.UpdateOauthClient(ctx, client)
}
