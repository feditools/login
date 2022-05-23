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

// CountApplicationTokens returns the number of application token.
func (c *Client) CountApplicationTokens(ctx context.Context) (int64, db.Error) {
	metric := c.metrics.NewDBQuery("CountApplicationTokens")

	count, err := c.newApplicationTokenQ((*models.ApplicationToken)(nil)).Count(ctx)
	if err != nil {
		go metric.Done(true)

		return 0, c.bun.errProc(err)
	}

	go metric.Done(false)

	return int64(count), nil
}

// CreateApplicationToken stores the application token.
func (c *Client) CreateApplicationToken(ctx context.Context, applicationToken *models.ApplicationToken) db.Error {
	metric := c.metrics.NewDBQuery("CreateApplicationToken")

	err := c.Create(ctx, applicationToken)
	if err != nil {
		go metric.Done(true)

		return c.bun.errProc(err)
	}

	go metric.Done(false)

	return nil
}

// ReadApplicationToken returns one application token.
func (c *Client) ReadApplicationToken(ctx context.Context, id int64) (*models.ApplicationToken, db.Error) {
	metric := c.metrics.NewDBQuery("ReadApplicationToken")

	applicationToken := new(models.ApplicationToken)
	err := c.newApplicationTokenQ(applicationToken).Where("id = ?", id).Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		go metric.Done(false)

		return nil, nil
	}
	if err != nil {
		go metric.Done(true)

		return nil, c.bun.ProcessError(err)
	}

	go metric.Done(false)

	return applicationToken, nil
}

// ReadApplicationTokenByToken returns one application token.
func (c *Client) ReadApplicationTokenByToken(ctx context.Context, token string) (*models.ApplicationToken, db.Error) {
	metric := c.metrics.NewDBQuery("ReadApplicationTokenByToken")

	applicationToken := &models.ApplicationToken{}

	err := c.newApplicationTokenQ(applicationToken).Where("token = ?", token).Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		go metric.Done(false)

		return nil, nil
	}
	if err != nil {
		go metric.Done(true)

		return nil, c.bun.ProcessError(err)
	}

	go metric.Done(false)

	return applicationToken, nil
}

// ReadApplicationTokensPage returns a page of application tokens.
func (c *Client) ReadApplicationTokensPage(ctx context.Context, index, count int) ([]*models.ApplicationToken, db.Error) {
	metric := c.metrics.NewDBQuery("ReadApplicationTokensPage")

	var applicationTokens []*models.ApplicationToken
	err := c.newApplicationTokensQ(&applicationTokens).
		Limit(count).
		Offset(libdatabase.Offset(index, count)).
		Scan(ctx)
	if err != nil {
		go metric.Done(true)

		return nil, c.bun.ProcessError(err)
	}

	go metric.Done(false)

	return applicationTokens, nil
}

// UpdateApplicationToken updates the stored application token.
func (c *Client) UpdateApplicationToken(ctx context.Context, client *models.ApplicationToken) db.Error {
	metric := c.metrics.NewDBQuery("UpdateApplicationToken")

	if err := c.Update(ctx, client); err != nil {
		go metric.Done(true)

		return c.bun.errProc(err)
	}

	go metric.Done(false)

	return nil
}

func (c *Client) newApplicationTokenQ(client *models.ApplicationToken) *bun.SelectQuery {
	return c.bun.
		NewSelect().
		Model(client)
}

func (c *Client) newApplicationTokensQ(clients *[]*models.ApplicationToken) *bun.SelectQuery {
	return c.bun.
		NewSelect().
		Model(clients)
}
