package webapp

import (
	"context"
	"github.com/feditools/login/internal/db"
	"github.com/feditools/login/internal/kv"
	"github.com/feditools/login/internal/kv/redis"
	"github.com/feditools/login/internal/token"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	oredis "github.com/go-oauth2/redis/v4"
	nethttp "net/http"
)

// OauthAuthorizeGetHandler handles oauth authorization
func (m *Module) OauthAuthorizeGetHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	_, shouldReturn := m.authRequireLoggedIn(w, r)
	if shouldReturn {
		return
	}

	err := m.oauth.HandleAuthorizeRequest(w, r)
	if err != nil {
		m.returnErrorPage(w, r, nethttp.StatusBadRequest, err.Error())
	}
}

// OauthTokenHandler handles oauth tokens
func (m *Module) OauthTokenHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	m.oauth.HandleTokenRequest(w, r)
}

func createOAuth(_ context.Context, d db.DB, r *redis.Client, t *token.Tokenizer) (*server.Server, error) {
	l := logger.WithField("func", "createOAuth")

	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	manager.MapTokenStorage(oredis.NewRedisStoreWithCli(
		r.RedisClient(),
		kv.KeyOauthToken(),
	))
	manager.MapClientStorage(NewAdapterClientStore(d, t))

	oauthServer := server.NewDefaultServer(manager)
	oauthServer.SetAllowGetAccessRequest(true)
	oauthServer.SetClientInfoHandler(server.ClientFormHandler)
	oauthServer.SetInternalErrorHandler(func(err error) *errors.Response {
		l.Errorf("Internal Error: %s", err.Error())
		return nil
	})
	oauthServer.SetResponseErrorHandler(func(re *errors.Response) {
		l.Errorf("Response Error: %s", re.Error.Error())
	})

	return oauthServer, nil
}
