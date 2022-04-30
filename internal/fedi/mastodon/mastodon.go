package mastodon

import (
	"github.com/feditools/login/internal/config"
	"github.com/feditools/login/internal/db"
	"github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/kv"
	"github.com/feditools/login/internal/models"
	"github.com/feditools/login/internal/token"
	mastodon "github.com/mattn/go-mastodon"
	"github.com/spf13/viper"
	"golang.org/x/sync/singleflight"
)

// Helper is a mastodon helper
type Helper struct {
	db   db.DB
	kv   kv.KV
	tokz *token.Tokenizer

	appClientName    string
	appWebsite       string
	externalHostname string

	registerAppGroup singleflight.Group
}

// New returns a new mastodon helper
func New(d db.DB, k kv.KV, t *token.Tokenizer, website string) (*Helper, error) {
	return &Helper{
		db:   d,
		kv:   k,
		tokz: t,

		appClientName:    viper.GetString(config.Keys.ApplicationName),
		appWebsite:       website,
		externalHostname: viper.GetString(config.Keys.ServerExternalHostname),
	}, nil
}

// newClient return new mastodon API client.
func newClient(instance *models.FediInstance, accessToken string) (*mastodon.Client, error) {
	l := logger.WithField("func", "newClient")

	// decrypt secret
	clientSecret, err := instance.GetClientSecret()
	if err != nil {
		l.Errorf("decrypting client secret: %s", err.Error())
		return nil, err
	}

	// create client
	client := mastodon.NewClient(&mastodon.Config{
		Server:       "https://" + instance.Domain,
		ClientID:     instance.ClientID,
		ClientSecret: clientSecret,
		AccessToken:  accessToken,
	})

	// apply custom transport
	client.Transport = &http.Transport{}
	return client, nil
}