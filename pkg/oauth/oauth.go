package oauth

import (
	"context"
	"encoding/gob"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

// Client is an oauth client.
type Client struct {
	config   oauth2.Config
	verifier *oidc.IDTokenVerifier
}

// New creates a new client
func New(ctx context.Context, cfg *Config) (*Client, error) {
	provider, err := oidc.NewProvider(ctx, cfg.ServerURL)
	if err != nil {
		return nil, err
	}

	c := oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		Scopes:       []string{oidc.ScopeOpenID},
		RedirectURL:  cfg.CallbackURL,
		Endpoint:     provider.Endpoint(),
	}

	oidcConfig := &oidc.Config{
		ClientID: cfg.ClientID,
	}
	v := provider.Verifier(oidcConfig)

	gob.Register(SessionKey(0))

	return &Client{
		config:   c,
		verifier: v,
	}, nil
}
