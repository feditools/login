package mastodon

import (
	"context"
	"github.com/feditools/login/internal/models"
	"net/url"
)

// GetAccessToken gets an access token for a account from a returned code
func (h *Helper) GetAccessToken(ctx context.Context, instance *models.FediInstance, code string) (accessToken string, err error) {
	// decrypt secret
	c, err := newClient(instance, "")
	if err != nil {
		return "", err
	}

	// authenticate
	instanceToken := h.tokz.GetToken(instance)
	err = c.AuthenticateToken(ctx, code, "https://"+h.externalHostname+"/callback/oauth/"+instanceToken)
	return c.Config.AccessToken, nil
}

// MakeLoginURL creates a login redirect url for mastodon
func (h *Helper) MakeLoginURL(_ context.Context, instance *models.FediInstance) (*url.URL, error) {
	instanceToken := h.tokz.GetToken(instance)
	u := &url.URL{
		Scheme: "https",
		Host:   instance.Domain,
		Path:   "/oauth/authorize",
	}
	q := u.Query()
	q.Set("client_id", instance.ClientID)
	q.Set("redirect_uri", "https://"+h.externalHostname+"/callback/oauth/"+instanceToken)
	q.Set("response_type", "code")
	q.Set("scope", "read:accounts")
	u.RawQuery = q.Encode()
	return u, nil
}
