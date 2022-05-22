package webapp

import (
	"context"
	"encoding/gob"
	htmltemplate "html/template"
	"net/url"
	"strings"
	"sync"

	"github.com/feditools/login/internal/oauth"

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
	"github.com/gorilla/sessions"
	"github.com/rbcervilla/redisstore/v8"
	"github.com/spf13/viper"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
)

// Module contains a webapp module for the web server. Implements web.Module.
type Module struct {
	db        db.DB
	fedi      *fedi.Fedi
	oauth     *oauth.Server
	store     sessions.Store
	language  *language.Module
	metrics   metrics.Collector
	minify    *minify.M
	srv       *http.Server
	templates *htmltemplate.Template
	tokenizer *token.Tokenizer

	logoSrcDark   string
	logoSrcLight  string
	headLinks     []libtemplate.HeadLink
	footerScripts []libtemplate.Script

	sigCache     map[string]string
	sigCacheLock sync.RWMutex
}

// New returns a new webapp module.
func New(ctx context.Context, d db.DB, r *redis.Client, f *fedi.Fedi, lMod *language.Module, oauthServer *oauth.Server, t *token.Tokenizer, mc metrics.Collector) (*Module, error) {
	l := logger.WithField("func", "New")

	// Fetch new store.
	store, err := redisstore.NewRedisStore(ctx, r.RedisClient())
	if err != nil {
		l.Errorf("create redis store: %s", err.Error())

		return nil, err
	}

	// parse external url
	externalURL, err := url.Parse(viper.GetString(config.Keys.ServerExternalURL))
	if err != nil {
		l.Errorf("parsing external url: %s", err.Error())

		return nil, err
	}

	store.KeyPrefix(kv.KeySession())
	store.Options(sessions.Options{
		Path:   "/",
		Domain: externalURL.Host,
		MaxAge: 86400 * 60,
	})

	// Register models for GOB
	gob.Register(http.SessionKey(0))
	gob.Register(models.FediAccount{})

	// minify
	var m *minify.M
	if viper.GetBool(config.Keys.ServerMinifyHTML) {
		m = minify.New()
		m.AddFunc("text/html", html.Minify)
	}

	// oauth
	oauthServer.SetUserAuthorizationHandler(oauthUserAuthorizeHandler)

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
			CrossOrigin: XOriginAnonymous,
			Integrity:   viper.GetString(config.Keys.WebappBootstrapCSSIntegrity),
		},
		{
			HRef:        viper.GetString(config.Keys.WebappFontAwesomeCSSURI),
			Rel:         "stylesheet",
			CrossOrigin: XOriginAnonymous,
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
			CrossOrigin: XOriginAnonymous,
			Integrity:   signature,
		})
	}

	// generate head links
	fs := []libtemplate.Script{
		{
			Src:         viper.GetString(config.Keys.WebappBootstrapJSURI),
			CrossOrigin: XOriginAnonymous,
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

// Name return the module name.
func (*Module) Name() string {
	return config.ServerRoleWebapp
}

// SetServer adds a reference to the server to the module.
func (m *Module) SetServer(s *http.Server) {
	m.srv = s
}
