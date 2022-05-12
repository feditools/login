package mastodon

import (
	"context"
	"errors"
	"fmt"
	"github.com/feditools/login/internal/fedi"
	"github.com/feditools/login/internal/models"
	"time"
)

// GetCurrentAccount retrieves the current federated account
func (h *Helper) GetCurrentAccount(ctx context.Context, instance *models.FediInstance, accessToken string) (*models.FediAccount, error) {
	l := logger.WithField("func", "GetCurrentAccount")

	// create mastodon client
	client, err := newClient(instance, accessToken)
	if err != nil {
		l.Errorf("creating client: %s", err.Error())
		return nil, err
	}

	// retrieve current account from
	account, err := client.GetAccountCurrentUser(ctx)
	if err != nil {
		l.Errorf("getting current account: %s", err.Error())
		return nil, err
	}

	// check if account is locked
	if account.Locked {
		return nil, fmt.Errorf("account '@%s@%s' locked", account.Username, instance.Domain)
	}

	// check if account is a bot
	if account.Bot {
		return nil, fmt.Errorf("account '@%s@%s' is a bot", account.Username, instance.Domain)
	}

	// check if account has moved
	if account.Moved != nil {
		return nil, fmt.Errorf("account '@%s@%s' has moved to '@%s'", account.Username, instance.Domain, account.Moved.Acct)
	}

	// try to retrieve federated account
	fediAccount, err := h.db.ReadFediAccountByUsername(ctx, instance.ID, account.Username)
	if err != nil {
		l.Errorf("db read: %s", err.Error())
		return nil, err
	}
	if fediAccount != nil {
		return fediAccount, nil
	}

	// do webfinger
	webFinger, err := h.fedi.GetWellknownWebFinger(ctx, account.Username, instance.Domain)
	if err != nil {
		l.Debugf("webfinger %s@%s: %s", account.Username, instance.Domain, err.Error())
		return nil, err
	}
	actorURI, err := fedi.FindActorURI(webFinger)
	if err != nil {
		l.Debugf("webfinger %s@%s: %s", account.Username, instance.Domain, err.Error())
		return nil, err
	}
	if actorURI == nil {
		msg := fmt.Sprintf("can't find actor uri for %s@%s", account.Username, instance.Domain)
		l.Debug(msg)
		return nil, errors.New(msg)
	}

	// create new federated account
	newFediAccount := &models.FediAccount{
		InstanceID:  instance.ID,
		Instance:    instance,
		ActorURI:    actorURI.String(),
		Username:    account.Username,
		DisplayName: account.DisplayName,
		LastFinger:  time.Now(),
	}
	err = newFediAccount.SetAccessToken(accessToken)
	if err != nil {
		l.Errorf("set access token: %s", err.Error())
		return nil, err
	}

	// write new federated account to database
	err = h.db.CreateFediAccount(ctx, newFediAccount)
	if err != nil {
		l.Errorf("db create: %s", err.Error())
		return nil, err
	}

	return newFediAccount, nil
}
