package webapp

import (
	"context"
	"encoding/gob"
	"github.com/feditools/login/internal/config"
	"github.com/feditools/login/internal/db"
	"github.com/feditools/login/internal/fedi"
	"github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/kv"
	"github.com/feditools/login/internal/kv/redis"
	"github.com/feditools/login/internal/language"
	"github.com/feditools/login/internal/metrics"
	"github.com/feditools/login/internal/models"
	"github.com/feditools/login/internal/path"
	"github.com/feditools/login/internal/template"
	"github.com/feditools/login/internal/token"
	"github.com/gorilla/sessions"
	"github.com/rbcervilla/redisstore/v8"
	"github.com/spf13/viper"
	minify "github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
	htmltemplate "html/template"
	"sync"
)

// Module contains a webapp module for the web server. Implements web.Module
type Module struct {
	db        db.DB
	fedi      *fedi.Fedi
	store     sessions.Store
	language  *language.Module
	metrics   metrics.Collector
	minify    *minify.M
	templates *htmltemplate.Template
	tokenizer *token.Tokenizer

	logoURI       string
	headLinks     []template.HeadLink
	footerScripts []template.Script

	sigCache     map[string]string
	sigCacheLock sync.RWMutex
}

// New returns a new webapp module
func New(ctx context.Context, db db.DB, r *redis.Client, f *fedi.Fedi, lMod *language.Module, t *token.Tokenizer, mc metrics.Collector) (http.Module, error) {
	l := logger.WithField("func", "New")

	// Fetch new store.
	store, err := redisstore.NewRedisStore(ctx, r.RedisClient())
	if err != nil {
		l.Errorf("create redis store: %s", err.Error())
		return nil, err
	}

	store.KeyPrefix(kv.KeySession())
	store.Options(sessions.Options{
		Path:   "/",
		Domain: viper.GetString(config.Keys.ServerExternalHostname),
		MaxAge: 86400 * 60,
	})

	// Register models for GOB
	gob.Register(SessionKey(0))
	gob.Register(models.FediAccount{})

	// minify
	var m *minify.M
	if viper.GetBool(config.Keys.ServerMinifyHTML) {
		m = minify.New()
		m.AddFunc("text/html", html.Minify)
	}

	// get templates
	tmpl, err := template.New(t)
	if err != nil {
		l.Errorf("create temates: %s", err.Error())
		return nil, err
	}

	// generate head links
	hl := []template.HeadLink{
		{
			HRef:        viper.GetString(config.Keys.WebappBootstrapCSSURI),
			Rel:         "stylesheet",
			CrossOrigin: "anonymous",
			Integrity:   viper.GetString(config.Keys.WebappBootstrapCSSIntegrity),
		},
	}
	paths := []string{
		path.FileDefaultCSS,
	}
	for _, path := range paths {
		signature, err := getSignature(DirWeb + path)
		if err != nil {
			l.Errorf("getting signature for %s: %s", path, err.Error())
		}

		hl = append(hl, template.HeadLink{
			HRef:        path,
			Rel:         "stylesheet",
			CrossOrigin: "anonymous",
			Integrity:   signature,
		})
	}

	// generate head links
	fs := []template.Script{
		{
			Src:         viper.GetString(config.Keys.WebappBootstrapJSURI),
			CrossOrigin: "anonymous",
			Integrity:   viper.GetString(config.Keys.WebappBootstrapJSIntegrity),
		},
	}

	return &Module{
		db:        db,
		fedi:      f,
		store:     store,
		language:  lMod,
		metrics:   mc,
		minify:    m,
		templates: tmpl,
		tokenizer: t,

		logoURI:       viper.GetString(config.Keys.WebappLogoURI),
		headLinks:     hl,
		footerScripts: fs,

		sigCache: map[string]string{},
	}, nil
}

// Name return the module name
func (m *Module) Name() string {
	return config.ServerRoleWebapp
}
