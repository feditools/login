package bun

import (
	"context"
	"database/sql"
	"errors"

	libdatabase "github.com/feditools/go-lib/database"
	"github.com/feditools/login/internal/db"
	"github.com/feditools/login/internal/models"
	"github.com/uptrace/bun"
)

// CountOauthClients returns the number of oauth clients.
func (c *Client) CountOauthClients(ctx context.Context) (int64, db.Error) {
	metric := c.metrics.NewDBQuery("CreateOauthClient")

	count, err := c.newOauthClientQ((*models.OauthClient)(nil)).Count(ctx)
	if err != nil {
		go metric.Done(true)
		return 0, c.bun.errProc(err)
	}

	go metric.Done(false)
	return int64(count), nil
}

// CreateOauthClient stores the oauth client.
func (c *Client) CreateOauthClient(ctx context.Context, client *models.OauthClient) db.Error {
	metric := c.metrics.NewDBQuery("CreateOauthClient")

	if err := c.Create(ctx, client); err != nil {
		go metric.Done(true)
		return c.bun.errProc(err)
	}

	go metric.Done(false)
	return nil
}

// ReadOauthClient returns one federated social account.
func (c *Client) ReadOauthClient(ctx context.Context, id int64) (*models.OauthClient, db.Error) {
	metric := c.metrics.NewDBQuery("ReadOauthClient")

	oauthClient := new(models.OauthClient)
	err := c.newOauthClientQ(oauthClient).Where("id = ?", id).Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		go metric.Done(false)
		return nil, nil
	}
	if err != nil {
		go metric.Done(true)
		return nil, c.bun.ProcessError(err)
	}

	go metric.Done(false)
	return oauthClient, nil
}

// ReadOauthClientsPage returns a page of oauth clients.
func (c *Client) ReadOauthClientsPage(ctx context.Context, index, count int) ([]*models.OauthClient, db.Error) {
	metric := c.metrics.NewDBQuery("ReadOauthClientsPage")

	var clients []*models.OauthClient
	err := c.newOauthClientsQ(&clients).
		Limit(count).
		Offset(libdatabase.Offset(index, count)).
		Scan(ctx)
	if err != nil {
		go metric.Done(true)
		return nil, c.bun.ProcessError(err)
	}

	go metric.Done(false)
	return clients, nil
}

// UpdateOauthClient updates the stored oauth client.
func (c *Client) UpdateOauthClient(ctx context.Context, client *models.OauthClient) db.Error {
	metric := c.metrics.NewDBQuery("UpdateOauthClient")

	if err := c.Update(ctx, client); err != nil {
		go metric.Done(true)
		return c.bun.errProc(err)
	}

	go metric.Done(false)
	return nil
}

func (c *Client) newOauthClientQ(client *models.OauthClient) *bun.SelectQuery {
	return c.bun.
		NewSelect().
		Model(client)
}

func (c *Client) newOauthClientsQ(clients *[]*models.OauthClient) *bun.SelectQuery {
	return c.bun.
		NewSelect().
		Model(clients)
}
