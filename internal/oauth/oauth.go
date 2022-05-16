package oauth

import (
	"context"
	"github.com/feditools/login/internal/db"
	"github.com/feditools/login/internal/kv"
	"github.com/feditools/login/internal/kv/redis"
	"github.com/feditools/login/internal/token"
	oerrors "github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	oredis "github.com/go-oauth2/redis/v4"
)

type Module struct {
	db  db.DB
	srv *server.Server
}

func New(_ context.Context, d db.DB, r *redis.Client, t *token.Tokenizer) (*server.Server, error) {
	l := logger.WithField("func", "New")

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
	oauthServer.SetInternalErrorHandler(func(err error) *oerrors.Response {
		l.Errorf("Internal Error: %s", err.Error())
		return nil
	})
	oauthServer.SetResponseErrorHandler(func(re *oerrors.Response) {
		l.Errorf("Response Error: %s", re.Error.Error())
	})

	return oauthServer, nil
}
