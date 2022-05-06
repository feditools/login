package fedi

import (
	"github.com/feditools/login/internal/config"
	"github.com/feditools/login/internal/db"
	"github.com/feditools/login/internal/fedi/mastodon"
	"github.com/feditools/login/internal/kv"
	"github.com/feditools/login/internal/token"
	"github.com/spf13/viper"
	"golang.org/x/sync/singleflight"
	"time"
)

// Fedi is a module for working with federated social instances
type Fedi struct {
	db   db.DB
	kv   kv.KV
	tokz *token.Tokenizer

	helpers map[Software]Helper

	appClientName    string
	externalHostname string

	nodeinfoCacheExp time.Duration

	requestGroup singleflight.Group
}

// New creates a new fedi module
func New(d db.DB, k kv.KV, t *token.Tokenizer) (*Fedi, error) {
	newFedi := &Fedi{
		db:   d,
		kv:   k,
		tokz: t,

		helpers: map[Software]Helper{},

		appClientName:    viper.GetString(config.Keys.ApplicationName),
		externalHostname: viper.GetString(config.Keys.ServerExternalHostname),

		nodeinfoCacheExp: 60 * time.Minute,
	}

	// mastodon helper
	mastoHelper, err := mastodon.New(d, k, t, AppWebsite)
	if err != nil {
		return nil, err
	}
	newFedi.helpers[SoftwareMastodon] = mastoHelper

	return newFedi, nil
}
