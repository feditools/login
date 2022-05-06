package fedi

import (
	"context"
	"github.com/feditools/login/internal/models"
	"net/url"
)

// Helper interacts with a federated social instance
type Helper interface {
	GetAccessToken(ctx context.Context, instance *models.FediInstance, code string) (accessToken string, err error)
	GetCurrentAccount(ctx context.Context, instance *models.FediInstance, accessToken string) (user *models.FediAccount, err error)
	GetSoftware() Software
	RegisterApp(ctx context.Context, instance *models.FediInstance) (clientID string, clientSecret string, err error)
	MakeLoginURL(ctx context.Context, instance *models.FediInstance) (url *url.URL, err error)
}

// Helper returns a helper for a given software package
func (f *Fedi) Helper(s Software) Helper {
	return f.helpers[s]
}
