package oauth

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"strings"
	"time"

	"github.com/go-oauth2/oauth2/v4"
	oerrors "github.com/go-oauth2/oauth2/v4/errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// JWTAccessClaims jwt claims.
type JWTAccessClaims struct {
	jwt.StandardClaims
}

// Valid claims verification.
func (a *JWTAccessClaims) Valid() error {
	if time.Unix(a.ExpiresAt, 0).Before(time.Now()) {
		return oerrors.ErrInvalidAccessToken
	}
	return nil
}

// NewAccessGenerator creates a new access token generator.
func (s *Server) NewAccessGenerator(issuer string, method jwt.SigningMethod) (*AccessGenerator, error) {
	return &AccessGenerator{
		Issuer:       issuer,
		SignedKeyID:  s.GetECPublicKeyID(),
		SignedKey:    s.GetECPrivateKey(),
		SignedMethod: method,
	}, nil
}

// AccessGenerator generate the jwt access token.
type AccessGenerator struct {
	Issuer       string
	SignedKeyID  string
	SignedKey    *ecdsa.PrivateKey
	SignedMethod jwt.SigningMethod
}

// Token based on the UUID generated token.
func (a *AccessGenerator) Token(ctx context.Context, data *oauth2.GenerateBasic, isGenRefresh bool) (string, string, error) {
	l := logger.WithField("func", "Token")
	l.Debugf("Called: %+v", data)
	l.Debugf("Client: %+v", data.Client)
	l.Debugf("Token: %+v", data.TokenInfo.GetCodeChallengeMethod())

	claims := &JWTAccessClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    a.Issuer,
			Audience:  data.Client.GetID(),
			Subject:   data.UserID,
			ExpiresAt: data.TokenInfo.GetAccessCreateAt().Add(data.TokenInfo.GetAccessExpiresIn()).Unix(),
		},
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
	refresh := ""

	if isGenRefresh {
		t := uuid.NewSHA1(uuid.Must(uuid.NewRandom()), []byte(access)).String()
		refresh = base64.URLEncoding.EncodeToString([]byte(t))
		refresh = strings.ToUpper(strings.TrimRight(refresh, "="))
	}

	return access, refresh, nil
}
