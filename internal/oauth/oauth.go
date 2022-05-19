package oauth

import (
	"context"
	"crypto"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/feditools/login/internal/config"
	"github.com/spf13/viper"
	nethttp "net/http"
	"strings"
	"time"

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
		AccessTokenExp:    time.Hour * 8,
		RefreshTokenExp:   time.Hour * 24 * 7,
		IsGenerateRefresh: true,
	}

	// access generator
	externalURL := strings.TrimSuffix(viper.GetString(config.Keys.ServerExternalURL), "/")
	accessGenerator, err := newServer.NewAccessGenerator(externalURL, jwt.SigningMethodES256)
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
		oauth:    oauthServer,
	}, nil
}

// keys

// GetECPrivateKey returns an ecdsa.PrivateKey.
func (s *Server) GetECPrivateKey() *ecdsa.PrivateKey {
	return s.keychain.ecdsa
}

// GetECPrivateKeyPEM returns a PEM encoded version of the private key.
func (s *Server) GetECPrivateKeyPEM() ([]byte, error) {
	privateKeyBytes, err := x509.MarshalECPrivateKey(s.keychain.ecdsa)
	if err != nil {
		return nil, err
	}

	newPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "EC PRIVATE KEY",
			Bytes: privateKeyBytes,
		},
	)

	return newPem, nil
}

// GetECPublicKeyCurve returns the ecdsa curve type.
func (s *Server) GetECPublicKeyCurve() string {
	return s.keychain.ecdsa.PublicKey.Curve.Params().Name
}

// GetECPublicKeyX returns the bytes of ecdsa X.
func (s *Server) GetECPublicKeyX() []byte {
	return s.keychain.ecdsa.PublicKey.X.Bytes()
}

// GetECPublicKeyY returns the bytes of ecdsa Y.
func (s *Server) GetECPublicKeyY() []byte {
	return s.keychain.ecdsa.PublicKey.Y.Bytes()
}

// GetECPublicKey returns a crypto.PublicKey compatable version of the public key.
func (s *Server) GetECPublicKey() crypto.PublicKey {
	return s.keychain.ecdsa.Public()
}

// GetECPublicKeyID returns the generated.
func (s *Server) GetECPublicKeyID() string {
	return s.keychain.ecdsaKID
}

// handlers

// HandleAuthorizeRequest passes an authorize request to the oauth server.
func (s *Server) HandleAuthorizeRequest(w nethttp.ResponseWriter, r *nethttp.Request) error {
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
