package bun

import (
	"context"
	"database/sql"
	"github.com/feditools/login/internal/db"
	"github.com/feditools/login/internal/models"
	"github.com/uptrace/bun"
)

// CreateFediAccount stores the federated social account
func (c *Client) CreateFediAccount(ctx context.Context, account *models.FediAccount) db.Error {
	metric := c.metrics.NewDBQuery("CreateFediAccount")

	err := c.Create(ctx, account)
	if err != nil {
		go metric.Done(true)
		return c.bun.errProc(err)
	}

	go metric.Done(false)
	return nil
}

// ReadFediAccount returns one federated social account
func (c *Client) ReadFediAccount(ctx context.Context, id int64) (*models.FediAccount, db.Error) {
	metric := c.metrics.NewDBQuery("ReadFediAccount")

	fediAccount := new(models.FediAccount)
	err := c.newFediAccountQ(fediAccount).Where("id = ?", id).Scan(ctx)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		go metric.Done(true)
		return nil, c.bun.ProcessError(err)
	}

	go metric.Done(false)
	return fediAccount, nil
}

// ReadFediAccountByUsername returns one federated social account
func (c *Client) ReadFediAccountByUsername(ctx context.Context, instanceID int64, username string) (*models.FediAccount, db.Error) {
	metric := c.metrics.NewDBQuery("ReadFediAccountByUsername")

	fediAccount := new(models.FediAccount)
	err := c.newFediAccountQ(fediAccount).
		ColumnExpr("fedi_account.*").
		Join("RIGHT JOIN fedi_instances").
		JoinOn("fedi_account.instance_id = fedi_instances.id").
		Where("fedi_instances.id = ?", instanceID).
		Where("lower(fedi_account.username) = lower(?)", username).
		Scan(ctx)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		go metric.Done(true)
		return nil, c.bun.ProcessError(err)
	}

	go metric.Done(false)
	return fediAccount, nil
}

// UpdateFediAccount updates the stored federated social account
func (c *Client) UpdateFediAccount(ctx context.Context, account *models.FediAccount) db.Error {
	metric := c.metrics.NewDBQuery("UpdateFediAccount")

	err := c.Update(ctx, account)
	if err != nil {
		go metric.Done(true)
		return c.bun.errProc(err)
	}

	go metric.Done(false)
	return nil
}

func (c *Client) newFediAccountQ(account *models.FediAccount) *bun.SelectQuery {
	return c.bun.
		NewSelect().
		Model(account)
}
