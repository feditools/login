package webapp

import (
	"context"
	"errors"
	"github.com/feditools/login/internal/db"
	"github.com/feditools/login/internal/token"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/models"
)

// AdapterClientStore adapts our database interface to the client store interface
type AdapterClientStore struct {
	db   db.DB
	tokz *token.Tokenizer
}

// NewAdapterClientStore creates a new client store adapter
func NewAdapterClientStore(d db.DB, t *token.Tokenizer) *AdapterClientStore {
	return &AdapterClientStore{
		db:   d,
		tokz: t,
	}
}

// GetByID returns a client based on an ID
func (c *AdapterClientStore) GetByID(ctx context.Context, tok string) (oauth2.ClientInfo, error) {

	l := logger.WithField("func", "GetByID")
	l.Debugf("looking for client %s", tok)

	if tok == "" {
		l.Debug("token empty")
		return nil, nil
	}

	kind, id, err := c.tokz.DecodeToken(tok)
	if err != nil {
		l.Debugf("error decoding token: %s", err.Error())
		return nil, err
	}
	if kind != token.KindOauthClient {
		l.Debugf("invalid token kind: %s", kind.String())
		return nil, errors.New("invalid token kind")
	}

	dbClient, err := c.db.ReadOauthClient(ctx, id)
	if err != nil {
		l.Errorf("db read: %s", err.Error())
		return nil, err
	}
	if dbClient == nil {
		l.Debug("client not found")
		return nil, nil
	}

	clientSecret, err := dbClient.GetSecret()
	if err != nil {
		l.Warnf("decoding secret: %s", err.Error())
		return nil, err
	}

	accountToken, err := c.tokz.EncodeToken(token.KindFediAccount, dbClient.OwnerID)
	if err != nil {
		l.Warnf("generating account tokens: %s", err.Error())
		return nil, err
	}

	newClient := &models.Client{
		ID:     tok,
		Secret: clientSecret,
		Domain: dbClient.RedirectURI,
		UserID: accountToken,
	}

	l.Debugf("new client: %s", newClient)

	return newClient, nil
}
