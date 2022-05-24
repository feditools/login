package oauth

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"errors"
	"strings"
	"time"

	"github.com/feditools/login/internal/config"
	"github.com/spf13/viper"

	"github.com/feditools/login/internal/kv"

	"github.com/go-oauth2/oauth2/v4"
	oerrors "github.com/go-oauth2/oauth2/v4/errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// JWTAccessClaims jwt claims.
type JWTAccessClaims struct {
	jwt.StandardClaims

	Nonce string `json:"nonce,omitempty"`
}

// Valid claims verification.
func (a *JWTAccessClaims) Valid() error {
	if time.Unix(a.ExpiresAt, 0).Before(time.Now()) {
		return oerrors.ErrInvalidAccessToken
	}

	return nil
}

// NewAccessGenerator creates a new access token generator.
func (s *Server) NewAccessGenerator(k kv.KV, issuer string, method jwt.SigningMethod) (*AccessGenerator, error) {
	return &AccessGenerator{
		kv: k,

		Issuer:          issuer,
		RefreshTokenExp: viper.GetDuration(config.Keys.AccessExpiration),
		SignedKeyID:     s.GetECPublicKeyID(),
		SignedKey:       s.GetECPrivateKey(),
		SignedMethod:    method,
	}, nil
}

// AccessGenerator generate the jwt access token.
type AccessGenerator struct {
	kv kv.KV

	Issuer          string
	RefreshTokenExp time.Duration
	SignedKeyID     string
	SignedKey       *ecdsa.PrivateKey
	SignedMethod    jwt.SigningMethod
}

// Token based on the UUID generated token.
func (a *AccessGenerator) Token(ctx context.Context, data *oauth2.GenerateBasic, isGenRefresh bool) (string, string, error) {
	l := logger.WithField("func", "Token")
	l.Debugf("Called: %+v", data)
	l.Debugf("Form: %+v", data.Request.Form)
	l.Debugf("Client: %+v", data.Client)
	l.Debugf("Token: %+v", data.TokenInfo)
	l.Debugf("IsGenRefresh: %v", isGenRefresh)

	// get nonce
	var nonce string
	if isGenRefresh {
		var err error
		nonce, err = a.kv.GetOauthNonceLogin(ctx, data.TokenInfo.GetUserID())
		if err != nil {
			l.Errorf("getting oauth nonce login: %s %T", err.Error(), err)

			return "", "", err
		}
	} else {
		var err error
		nonce, err = a.kv.GetOauthNonceRefresh(ctx, data.TokenInfo.GetRefresh())
		if err != nil {
			l.Errorf("getting oauth nonce refresh: %s %T", err.Error(), err)

			return "", "", err
		}
	}
	if nonce == "" {
		msg := "missing oauth nonce"
		l.Error(msg)

		return "", "", errors.New(msg)
	}

	// build jwt
	claims := &JWTAccessClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    a.Issuer,
			Audience:  data.Client.GetID(),
			Subject:   data.UserID,
			ExpiresAt: data.TokenInfo.GetAccessCreateAt().Add(data.TokenInfo.GetAccessExpiresIn()).Unix(),
		},
		Nonce: nonce,
	}

	token := jwt.NewWithClaims(a.SignedMethod, claims)
	if a.SignedKeyID != "" {
		token.Header["kid"] = a.SignedKeyID
	}

	access, err := token.SignedString(a.SignedKey)
	if err != nil {
		l.Errorf("signing string: %s", err.Error())

		return "", "", err
	}

	// generate refresh token if asked
	refresh := ""
	if isGenRefresh {
		t := uuid.NewSHA1(uuid.Must(uuid.NewRandom()), []byte(access)).String()
		refresh = base64.URLEncoding.EncodeToString([]byte(t))
		refresh = strings.ToUpper(strings.TrimRight(refresh, "="))

		// move nonce
		err = a.kv.DeleteOauthNonceLogin(ctx, data.TokenInfo.GetUserID())
		if err != nil {
			l.Errorf("deleting oauth nonce login: %s %T", err.Error(), err)

			return "", "", err
		}
		err = a.kv.SetOauthNonceRefresh(ctx, refresh, nonce, a.RefreshTokenExp)
		if err != nil {
			l.Errorf("set oauth nonce refresh: %s", err.Error())

			return "", "", err
		}
	}

	return access, refresh, nil
}
