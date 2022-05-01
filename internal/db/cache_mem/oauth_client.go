package cachemem

import (
	"context"
	"github.com/feditools/login/internal/db"
	"github.com/feditools/login/internal/models"
)

// ReadOauthClient returns one oauth client
func (c *CacheMem) ReadOauthClient(ctx context.Context, id int64) (*models.OauthClient, db.Error) {
	return c.db.ReadOauthClient(ctx, id)
}
