package cachemem

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/gob"
	"errors"
	"strconv"

	"github.com/allegro/bigcache/v3"
	"github.com/feditools/login/internal/db"
	"github.com/feditools/login/internal/models"
)

// CountApplicationTokens returns the number of application token.
func (c *CacheMem) CountApplicationTokens(ctx context.Context) (int64, db.Error) {
	metric := c.metrics.NewDBCacheQuery("CountApplicationTokens")

	count, hit := c.getCount(ctx, keyCountApplicationTokens())
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
		c.setCount(ctx, keyCountApplicationTokens(), count)
	}
	go metric.Done(false, false)

	return count, nil
}

// CreateApplicationToken stores the application token.
func (c *CacheMem) CreateApplicationToken(ctx context.Context, applicationToken *models.ApplicationToken) db.Error {
	err := c.db.CreateApplicationToken(ctx, applicationToken)
	if err != nil {
		return err
	}
	c.setApplicationToken(ctx, applicationToken)

	return nil
}

// ReadApplicationToken returns one application token.
func (c *CacheMem) ReadApplicationToken(ctx context.Context, id int64) (*models.ApplicationToken, db.Error) {
	metric := c.metrics.NewDBCacheQuery("ReadApplicationToken")

	applicationToken, hit := c.getApplicationToken(ctx, id)
	if hit {
		go metric.Done(true, false)

		return applicationToken, nil
	}
	applicationToken, err := c.db.ReadApplicationToken(ctx, id)
	if err != nil {
		go metric.Done(false, true)

		return nil, err
	}
	if applicationToken != nil {
		c.setApplicationToken(ctx, applicationToken)
	}
	go metric.Done(false, false)

	return applicationToken, nil
}

// ReadApplicationTokenByToken returns one application token.
func (c *CacheMem) ReadApplicationTokenByToken(ctx context.Context, token string) (*models.ApplicationToken, db.Error) {
	metric := c.metrics.NewDBCacheQuery("ReadApplicationTokenByToken")

	applicationToken, hit := c.getApplicationTokenByToken(ctx, token)
	if hit {
		go metric.Done(true, false)

		return applicationToken, nil
	}
	applicationToken, err := c.db.ReadApplicationTokenByToken(ctx, token)
	if err != nil {
		go metric.Done(false, true)

		return nil, err
	}
	if applicationToken != nil {
		c.setApplicationToken(ctx, applicationToken)
	}

	go metric.Done(false, false)

	return applicationToken, nil
}

// ReadApplicationTokensPage returns a page of application tokens.
func (c *CacheMem) ReadApplicationTokensPage(ctx context.Context, index, count int) ([]*models.ApplicationToken, db.Error) {
	return c.db.ReadApplicationTokensPage(ctx, index, count)
}

// UpdateApplicationToken updates the stored application token.
func (c *CacheMem) UpdateApplicationToken(ctx context.Context, applicationToken *models.ApplicationToken) db.Error {
	err := c.db.UpdateApplicationToken(ctx, applicationToken)
	if err != nil {
		return err
	}
	c.setApplicationToken(ctx, applicationToken)

	return nil
}

func (c *CacheMem) getApplicationToken(_ context.Context, id int64) (*models.ApplicationToken, bool) {
	l := logger.WithField("func", "getApplicationToken")

	// check cache
	entry, err := c.applicationToken.Get(strconv.FormatInt(id, 10))
	if errors.Is(err, bigcache.ErrEntryNotFound) {
		return nil, false
	}
	if err != nil {
		l.Warnf("cache get: %s", err.Error())

		return nil, false
	}

	// decode
	buf := bytes.NewBuffer(entry)
	dec := gob.NewDecoder(buf)
	applicationToken := new(models.ApplicationToken)
	if err := dec.Decode(&applicationToken); err != nil {
		l.Warnf("cache decode: %s", err.Error())

		return nil, false
	}

	return applicationToken, true
}

func (c *CacheMem) getApplicationTokenByToken(ctx context.Context, token string) (*models.ApplicationToken, bool) {
	l := logger.WithField("func", "getApplicationTokenByToken")

	// check username cache
	entry, err := c.applicationTokenTokenToID.Get(token)
	if errors.Is(err, bigcache.ErrEntryNotFound) {
		return nil, false
	}
	if err != nil {
		l.Warnf("cache get: %s", err.Error())

		return nil, false
	}
	i := int64(binary.LittleEndian.Uint64(entry))

	return c.getApplicationToken(ctx, i)
}

func (c *CacheMem) setApplicationToken(_ context.Context, applicationToken *models.ApplicationToken) {
	l := logger.WithField("func", "setApplicationToken")

	// encode object
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(applicationToken); err != nil {
		l.Warnf("cache encode: %s", err.Error())

		return
	}

	// put in the cache
	err := c.applicationToken.Set(strconv.FormatInt(applicationToken.ID, 10), buf.Bytes())
	if err != nil {
		l.Warnf("cache obj: %s", err.Error())

		return
	}

	// encode id
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(applicationToken.ID))
	err = c.applicationTokenTokenToID.Set(applicationToken.Token, b)
	if err != nil {
		l.Warnf("cache token: %s", err.Error())

		return
	}
}
