package cachemem

import (
	"context"
	"github.com/feditools/login/internal/db"
	"github.com/feditools/login/internal/models"
)

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
