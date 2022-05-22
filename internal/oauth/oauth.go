package oauth

import (
	"context"
	"crypto/ecdsa"
	"errors"
	nethttp "net/http"
	"strconv"
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

const (
	accessTokenExp  = time.Hour * 8
	refreshTokenExp = time.Hour * 24 * 7
)

// Server is an oauth server.
type Server struct {
	keychain *Keychain
	kv       kv.KV
	oauth    *server.Server
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
		AccessTokenExp:    accessTokenExp,
		RefreshTokenExp:   refreshTokenExp,
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

// handlers

// HandleAuthorizeRequest passes an authorize request to the oauth server.
func (s *Server) HandleAuthorizeRequest(uid int64, w nethttp.ResponseWriter, r *nethttp.Request) error {
	l := logger.WithField("func", "HandleAuthorizeRequest")

	nonce := r.Form.Get("nonce")
	l.Debugf("nonce: '%s'", nonce)
	if nonce == "" {
		return errors.New("missing nonce")
	}

	sessionID := r.Form.Get("session_id")
	l.Debugf("nonce: '%s'", sessionID)
	if sessionID == "" {
		return errors.New("missing session id")
	}

	err := s.kv.SetOauthNonce(r.Context(), strconv.FormatInt(uid, 10), sessionID, nonce, refreshTokenExp*2)
	if err != nil {
		l.Errorf("set oauth nonce: %s", err.Error())

		return err
	}

	return s.oauth.HandleAuthorizeRequest(w, r)
}

// HandleTokenRequest passes a token request to the oauth server.
func (s *Server) HandleTokenRequest(w nethttp.ResponseWriter, r *nethttp.Request) error {
	return s.oauth.HandleTokenRequest(w, r)
}

// SetUserAuthorizationHandler sets the UserAuthorizationHandler on the OAuth server.
func (s *Server) SetUserAuthorizationHandler(h server.UserAuthorizationHandler) {
	s.oauth.UserAuthorizationHandler = h
}
