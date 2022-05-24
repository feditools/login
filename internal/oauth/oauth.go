package oauth

import (
	"context"
	"crypto/ecdsa"
	"strings"
	"time"

	"github.com/feditools/login/internal/config"
	"github.com/spf13/viper"

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

// Server is an oauth server.
type Server struct {
	keychain *Keychain
	kv       kv.KV
	oauth    *server.Server

	LoginNonceExp time.Duration
}

// New returns a new oauth server.Server object.
func New(_ context.Context, d db.DB, r *redis.Client, t *token.Tokenizer) (*Server, error) {
	// l := logger.WithField("func", "New")

	// read keys
	keychain, err := NewKeychain()
	if err != nil {
		return nil, err
	}

	// create server
	newServer := &Server{
		keychain: keychain,
	}

	// authorize token config
	authorizeCodeTokenCfg := &manage.Config{
		AccessTokenExp:    viper.GetDuration(config.Keys.AccessExpiration),
		RefreshTokenExp:   viper.GetDuration(config.Keys.RefreshExpiration),
		IsGenerateRefresh: true,
	}

	// access generator
	externalURL := strings.TrimSuffix(viper.GetString(config.Keys.ServerExternalURL), "/")
	accessGenerator, err := newServer.NewAccessGenerator(r, externalURL, jwt.SigningMethodES256)
	if err != nil {
		return nil, err
	}

	// create oauth manager
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(authorizeCodeTokenCfg)
	manager.MapAccessGenerate(accessGenerator)
	// manager.MapAccessGenerate(generates.NewAccessGenerate())
	manager.MapTokenStorage(oredis.NewRedisStoreWithCli(
		r.RedisClient(),
		kv.KeyOauthToken(),
	))
	manager.MapClientStorage(NewAdapterClientStore(d, t))

	// create oauth server
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

	return &Server{
		keychain: keychain,
		kv:       r,
		oauth:    oauthServer,

		LoginNonceExp: viper.GetDuration(config.Keys.LoginNonceExpiration),
	}, nil
}

// keys

// GetECPrivateKey returns an ecdsa.PrivateKey.
func (s *Server) GetECPrivateKey() *ecdsa.PrivateKey {
	return s.keychain.ecdsa
}

// GetECPublicKey returns a crypto.PublicKey compatable version of the public key.
func (s *Server) GetECPublicKey() *ecdsa.PublicKey {
	pubKey := s.keychain.ecdsa.PublicKey

	return &pubKey
}

// GetECPublicKeyID returns the generated.
func (s *Server) GetECPublicKeyID() string {
	return s.keychain.ecdsaKID
}

// SetUserAuthorizationHandler sets the UserAuthorizationHandler on the OAuth server.
func (s *Server) SetUserAuthorizationHandler(h server.UserAuthorizationHandler) {
	s.oauth.UserAuthorizationHandler = h
}
