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

// CountFediInstances returns the number of federated instance.
func (c *Client) CountFediInstances(ctx context.Context) (int64, db.Error) {
	metric := c.metrics.NewDBQuery("CountFediInstances")

	count, err := c.newFediInstanceQ((*models.FediInstance)(nil)).Count(ctx)
	if err != nil {
		go metric.Done(true)

		return 0, c.bun.errProc(err)
	}

	go metric.Done(false)

	return int64(count), nil
}

// CreateFediInstance stores the federated instance.
func (c *Client) CreateFediInstance(ctx context.Context, instance *models.FediInstance) db.Error {
	metric := c.metrics.NewDBQuery("CreateFediInstance")

	if err := c.Create(ctx, instance); err != nil {
		go metric.Done(true)

		return c.bun.errProc(err)
	}

	go metric.Done(false)

	return nil
}

// ReadFediInstance returns one federated social instance.
func (c *Client) ReadFediInstance(ctx context.Context, id int64) (*models.FediInstance, db.Error) {
	metric := c.metrics.NewDBQuery("ReadFediInstance")

	fediInstance := &models.FediInstance{}

	err := c.newFediInstanceQ(fediInstance).Where("id = ?", id).Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		go metric.Done(true)

		return nil, c.bun.ProcessError(err)
	}

	go metric.Done(false)

	return fediInstance, nil
}

// ReadFediInstanceByDomain returns one federated social instance.
func (c *Client) ReadFediInstanceByDomain(ctx context.Context, domain string) (*models.FediInstance, db.Error) {
	metric := c.metrics.NewDBQuery("ReadFediInstanceByDomain")

	fediInstance := &models.FediInstance{}

	err := c.newFediInstanceQ(fediInstance).Where("lower(domain) = lower(?)", domain).Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		go metric.Done(true)

		return nil, c.bun.ProcessError(err)
	}

	go metric.Done(false)

	return fediInstance, nil
}

// ReadFediInstancesPage returns a page of federated social instances.
func (c *Client) ReadFediInstancesPage(ctx context.Context, index, count int) ([]*models.FediInstance, db.Error) {
	metric := c.metrics.NewDBQuery("ReadFediInstancesPage")

	var instances []*models.FediInstance

	err := c.newFediInstancesQ(&instances).
		Limit(count).
		Offset(libdatabase.Offset(index, count)).
		Scan(ctx)
	if err != nil {
		go metric.Done(true)

		return nil, c.bun.ProcessError(err)
	}

	go metric.Done(false)

	return instances, nil
}

// UpdateFediInstance updates the stored federated instance.
func (c *Client) UpdateFediInstance(ctx context.Context, instance *models.FediInstance) db.Error {
	metric := c.metrics.NewDBQuery("UpdateFediInstance")

	if err := c.Update(ctx, instance); err != nil {
		go metric.Done(true)

		return c.bun.errProc(err)
	}

	go metric.Done(false)

	return nil
}

func (c *Client) newFediInstanceQ(instance *models.FediInstance) *bun.SelectQuery {
	return c.bun.
		NewSelect().
		Model(instance)
}

func (c *Client) newFediInstancesQ(instances *[]*models.FediInstance) *bun.SelectQuery {
	return c.bun.
		NewSelect().
		Model(instances)
}
