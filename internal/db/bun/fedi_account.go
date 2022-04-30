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
	return c.Create(ctx, account)
}

// ReadFediAccount returns one federated social account
func (c *Client) ReadFediAccount(ctx context.Context, id int64) (*models.FediAccount, db.Error) {
	fediAccount := new(models.FediAccount)
	err := c.newFediAccountQ(fediAccount).Where("id = ?", id).Scan(ctx)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, c.bun.ProcessError(err)
	}
	return fediAccount, nil
}

// ReadFediAccountByUsername returns one federated social account
func (c *Client) ReadFediAccountByUsername(ctx context.Context, instanceID int64, username string) (*models.FediAccount, db.Error) {
	var account *models.FediAccount
	err := c.newFediAccountQ(account).
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
		return nil, c.bun.ProcessError(err)
	}
	return account, nil
}

// UpdateFediAccount updates the stored federated social account
func (c *Client) UpdateFediAccount(ctx context.Context, account *models.FediAccount) db.Error {
	return c.Create(ctx, account)
}

func (c *Client) newFediAccountQ(account *models.FediAccount) *bun.SelectQuery {
	return c.bun.
		NewSelect().
		Model(account)
}
