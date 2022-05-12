package webapp

import (
	"context"
	"encoding/gob"
	"github.com/feditools/go-lib/language"
	"github.com/feditools/go-lib/metrics"
	libtemplate "github.com/feditools/go-lib/template"
	"github.com/feditools/login/internal/config"
	"github.com/feditools/login/internal/db"
	"github.com/feditools/login/internal/fedi"
	"github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/http/template"
	"github.com/feditools/login/internal/kv"
	"github.com/feditools/login/internal/kv/redis"
	"github.com/feditools/login/internal/models"
	"github.com/feditools/login/internal/path"
	"github.com/feditools/login/internal/token"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/gorilla/sessions"
	"github.com/rbcervilla/redisstore/v8"
	"github.com/spf13/viper"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
	htmltemplate "html/template"
	"strings"
	"sync"
)

// Module contains a webapp module for the web server. Implements web.Module
type Module struct {
	db        db.DB
	fedi      *fedi.Fedi
	oauth     *server.Server
	store     sessions.Store
	language  *language.Module
	metrics   metrics.Collector
	minify    *minify.M
	templates *htmltemplate.Template
	tokenizer *token.Tokenizer

	logoSrcDark   string
	logoSrcLight  string
	headLinks     []libtemplate.HeadLink
	footerScripts []libtemplate.Script

	sigCache     map[string]string
	sigCacheLock sync.RWMutex
}

// New returns a new webapp module
func New(ctx context.Context, d db.DB, r *redis.Client, f *fedi.Fedi, lMod *language.Module, t *token.Tokenizer, mc metrics.Collector) (http.Module, error) {
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

	// oauth
	oauthServer, err := createOAuth(ctx, d, r, t)
	if err != nil {
		l.Errorf("create oauth server: %s", err.Error())
		return nil, err
	}

	// get templates
	tmpl, err := template.New(t)
	if err != nil {
		l.Errorf("create templates: %s", err.Error())
		return nil, err
	}

	// generate head links
	hl := []libtemplate.HeadLink{
		{
			HRef:        viper.GetString(config.Keys.WebappBootstrapCSSURI),
			Rel:         "stylesheet",
			CrossOrigin: "anonymous",
			Integrity:   viper.GetString(config.Keys.WebappBootstrapCSSIntegrity),
		},
		{
			HRef:        viper.GetString(config.Keys.WebappFontAwesomeCSSURI),
			Rel:         "stylesheet",
			CrossOrigin: "anonymous",
			Integrity:   viper.GetString(config.Keys.WebappFontAwesomeCSSIntegrity),
		},
	}
	paths := []string{
		path.FileDefaultCSS,
	}
	for _, filePath := range paths {
		signature, err := getSignature(strings.TrimPrefix(filePath, "/"))
		if err != nil {
			l.Errorf("getting signature for %s: %s", filePath, err.Error())
		}

		hl = append(hl, libtemplate.HeadLink{
			HRef:        filePath,
			Rel:         "stylesheet",
			CrossOrigin: "anonymous",
			Integrity:   signature,
		})
	}

	// generate head links
	fs := []libtemplate.Script{
		{
			Src:         viper.GetString(config.Keys.WebappBootstrapJSURI),
			CrossOrigin: "anonymous",
			Integrity:   viper.GetString(config.Keys.WebappBootstrapJSIntegrity),
		},
	}

	return &Module{
		db:        d,
		fedi:      f,
		oauth:     oauthServer,
		store:     store,
		language:  lMod,
		metrics:   mc,
		minify:    m,
		templates: tmpl,
		tokenizer: t,

		logoSrcDark:   viper.GetString(config.Keys.WebappLogoSrcDark),
		logoSrcLight:  viper.GetString(config.Keys.WebappLogoSrcLight),
		headLinks:     hl,
		footerScripts: fs,

		sigCache: map[string]string{},
	}, nil
}

// Name return the module name
func (m *Module) Name() string {
	return config.ServerRoleWebapp
}
