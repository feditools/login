package token

import (
	"github.com/feditools/login/internal/config"
	"github.com/feditools/login/internal/models"
	"github.com/speps/go-hashids/v2"
	"github.com/spf13/viper"
)

// Tokenizer generates public tokens for database IDs to obfuscate the database IDs.
type Tokenizer struct {
	h *hashids.HashID
}

// New returns a new tokenizer.
func New() (*Tokenizer, error) {
	// set config
	hd := hashids.NewData()
	salt := viper.GetString(config.Keys.TokenSalt)
	if salt == "" {
		return nil, ErrSaltEmpty
	}
	hd.Salt = salt
	hd.MinLength = 16

	// create hashid
	hid, err := hashids.NewWithData(hd)
	if err != nil {
		return nil, err
	}

	return &Tokenizer{
		h: hid,
	}, nil
}

// DecodeToken returns the kind and id number of a provided token.
func (t *Tokenizer) DecodeToken(token string) (Kind, int64, error) {
	d, err := t.h.DecodeInt64WithError(token)
	if err != nil {
		return 0, 0, err
	}
	if len(d) != 2 {
		return 0, 0, ErrInvalidLength
	}
	return Kind(d[0]), d[1], nil
}

// EncodeToken turns a model kind and id into a token.
func (t *Tokenizer) EncodeToken(kind Kind, id int64) (string, error) {
	return t.h.EncodeInt64([]int64{int64(kind), id})
}

// GetToken returns a token for a known model type.
func (t *Tokenizer) GetToken(o interface{}) string {
	l := logger.WithField("func", "GetToken")

	switch o := o.(type) {
	case models.ApplicationToken:
		tok, err := t.EncodeToken(KindApplicationToken, o.ID)
		if err != nil {
			l.Errorf(errRespGenerateToken, KindApplicationToken, err.Error())
		}

		return tok
	case *models.ApplicationToken:
		tok, err := t.EncodeToken(KindApplicationToken, o.ID)
		if err != nil {
			l.Errorf(errRespGenerateToken, KindApplicationToken, err.Error())
		}

		return tok
	case models.FediAccount:
		tok, err := t.EncodeToken(KindFediAccount, o.ID)
		if err != nil {
			l.Errorf(errRespGenerateToken, KindFediAccount, err.Error())
		}

		return tok
	case *models.FediAccount:
		tok, err := t.EncodeToken(KindFediAccount, o.ID)
		if err != nil {
			l.Errorf(errRespGenerateToken, KindFediAccount, err.Error())
		}

		return tok
	case models.FediInstance:
		tok, err := t.EncodeToken(KindFediInstance, o.ID)
		if err != nil {
			l.Errorf(errRespGenerateToken, KindFediInstance, err.Error())
		}

		return tok
	case *models.FediInstance:
		tok, err := t.EncodeToken(KindFediInstance, o.ID)
		if err != nil {
			l.Errorf(errRespGenerateToken, KindFediInstance, err.Error())
		}

		return tok
	case models.OauthClient:
		tok, err := t.EncodeToken(KindOauthClient, o.ID)
		if err != nil {
			l.Errorf(errRespGenerateToken, KindOauthClient, err.Error())
		}

		return tok
	case *models.OauthClient:
		tok, err := t.EncodeToken(KindOauthClient, o.ID)
		if err != nil {
			l.Errorf(errRespGenerateToken, KindOauthClient, err.Error())
		}

		return tok
	case models.OauthScope:
		tok, err := t.EncodeToken(KindOauthScope, o.ID)
		if err != nil {
			l.Errorf(errRespGenerateToken, KindOauthScope, err.Error())
		}

		return tok
	case *models.OauthScope:
		tok, err := t.EncodeToken(KindOauthScope, o.ID)
		if err != nil {
			l.Errorf(errRespGenerateToken, KindOauthScope, err.Error())
		}

		return tok
	default:
		l.Errorf("unknown kind: %T", o)

		return ""
	}
}
