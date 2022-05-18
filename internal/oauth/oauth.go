package oauth

import (
	"context"

	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/golang-jwt/jwt"

	"github.com/feditools/login/internal/db"
	"github.com/feditools/login/internal/kv"
	"github.com/feditools/login/internal/kv/redis"
	"github.com/feditools/login/internal/token"
	oerrors "github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	oredis "github.com/go-oauth2/redis/v4"
)

// New returns a new oauth server.Server object.
func New(_ context.Context, d db.DB, r *redis.Client, t *token.Tokenizer) (*server.Server, error) {
	// l := logger.WithField("func", "New")

	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("00000000"), jwt.SigningMethodHS512))
	manager.MapTokenStorage(oredis.NewRedisStoreWithCli(
		r.RedisClient(),
		kv.KeyOauthToken(),
	))
	manager.MapClientStorage(NewAdapterClientStore(d, t))

	oauthServer := server.NewDefaultServer(manager)
	oauthServer.SetAllowGetAccessRequest(true)
	oauthServer.SetClientInfoHandler(server.ClientFormHandler)
	oauthServer.SetInternalErrorHandler(func(err error) *oerrors.Response {
		l := logger.WithField("func", "SetInternalErrorHandler")
		l.Errorf("Internal Error: %s", err.Error())
		return nil
	})
	oauthServer.SetResponseErrorHandler(func(re *oerrors.Response) {
		l := logger.WithField("func", "SetResponseErrorHandler")
		l.Errorf("Response Error: %s", re.Error.Error())
	})

	return oauthServer, nil
}
