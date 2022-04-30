package bun

import (
	"context"
	"database/sql"
	"github.com/feditools/login/internal/db"
	"github.com/feditools/login/internal/models"
	"github.com/uptrace/bun"
)

// ReadOauthClient returns one federated social account
func (c *Client) ReadOauthClient(ctx context.Context, id int64) (*models.OauthClient, db.Error) {
	oauthClient := new(models.OauthClient)
	err := c.newOauthClientQ(oauthClient).Where("id = ?", id).Scan(ctx)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, c.bun.ProcessError(err)
	}
	return oauthClient, nil
}

func (c *Client) newOauthClientQ(client *models.OauthClient) *bun.SelectQuery {
	return c.bun.
		NewSelect().
		Model(client)
}
