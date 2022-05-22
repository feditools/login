package oauth

import (
	"context"
	"encoding/gob"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type Client struct {
	config   oauth2.Config
	verifier *oidc.IDTokenVerifier
}

func New(ctx context.Context, callbackURL, serverURL, clientID, secret string) (*Client, error) {
	provider, err := oidc.NewProvider(ctx, serverURL)
	if err != nil {
		return nil, err
	}

	c := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: secret,
		Scopes:       []string{oidc.ScopeOpenID},
		RedirectURL:  callbackURL,
		Endpoint:     provider.Endpoint(),
	}

	oidcConfig := &oidc.Config{
		ClientID: clientID,
	}
	v := provider.Verifier(oidcConfig)

	gob.Register(SessionKey(0))

	return &Client{
		config:   c,
		verifier: v,
	}, nil
}
