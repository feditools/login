package bun

import (
	"context"
	"database/sql"
	"github.com/feditools/login/internal/db"
	"github.com/feditools/login/internal/models"
	"github.com/uptrace/bun"
)

// CreateFediInstance stores the federated instance
func (c *Client) CreateFediInstance(ctx context.Context, instance *models.FediInstance) db.Error {
	return c.Create(ctx, instance)
}

// ReadFediInstance returns one federated social instance
func (c *Client) ReadFediInstance(ctx context.Context, id int64) (*models.FediInstance, db.Error) {
	fediInstance := &models.FediInstance{}

	err := c.newFediInstanceQ(fediInstance).Where("id = ?", id).Scan(ctx)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, c.bun.ProcessError(err)
	}

	return fediInstance, nil
}

// ReadFediInstanceByDomain returns one federated social instance
func (c *Client) ReadFediInstanceByDomain(ctx context.Context, domain string) (*models.FediInstance, db.Error) {
	fediInstance := &models.FediInstance{}

	err := c.newFediInstanceQ(fediInstance).Where("lower(domain) = lower(?)", domain).Scan(ctx)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, c.bun.ProcessError(err)
	}

	return fediInstance, nil
}

// UpdateFediInstance updates the stored federated instance
func (c *Client) UpdateFediInstance(ctx context.Context, instance *models.FediInstance) db.Error {
	return c.Update(ctx, instance)
}

func (c *Client) newFediInstanceQ(instance *models.FediInstance) *bun.SelectQuery {
	return c.bun.
		NewSelect().
		Model(instance)
}
