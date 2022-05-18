package mastodon

import (
	"context"
	"net/http"

	fthttp "github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/models"
	mastodon "github.com/mattn/go-mastodon"
)

// RegisterApp registers feditools with mastodon and returns the client id and client secret.
func (h *Helper) RegisterApp(ctx context.Context, instance *models.FediInstance) (string, string, error) {
	l := logger.WithField("func", "RegisterApp")
	v, err, _ := h.registerAppGroup.Do(instance.Domain, func() (interface{}, error) {
		instanceToken := h.tokz.GetToken(instance)
		app, err := mastodon.RegisterApp(ctx, &mastodon.AppConfig{
			Client: http.Client{
				Transport: &fthttp.Transport{},
			},
			Server:       "https://" + instance.ServerHostname,
			ClientName:   h.appClientName,
			Scopes:       "read:accounts",
			Website:      h.appWebsite,
			RedirectURIs: h.externalURL + "/callback/oauth/" + instanceToken,
		})

		if err != nil {
			l.Errorf("registering app: %s", err.Error())
			return nil, err
		}

		keys := []string{
			app.ClientID,
			app.ClientSecret,
		}

		return &keys, nil
	})

	if err != nil {
		l.Errorf("singleflight: %s", err.Error())
		return "", "", err
	}

	keys := v.(*[]string)
	return (*keys)[0], (*keys)[1], nil
}
