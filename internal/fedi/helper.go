package fedi

import (
	"context"
	"net/url"

	"github.com/feditools/login/internal/models"
)

// Helper interacts with a federated social instance.
type Helper interface {
	GetAccessToken(ctx context.Context, instance *models.FediInstance, code string) (accessToken string, err error)
	GetCurrentAccount(ctx context.Context, instance *models.FediInstance, accessToken string) (user *models.FediAccount, err error)
	GetSoftware() Software
	RegisterApp(ctx context.Context, instance *models.FediInstance) (clientID string, clientSecret string, err error)
	SetFedi(f *Fedi)
	MakeLoginURI(ctx context.Context, instance *models.FediInstance) (loginURI *url.URL, err error)
}

// Helper returns a helper for a given software package.
func (f *Fedi) Helper(s Software) Helper {
	return f.helpers[s]
}
