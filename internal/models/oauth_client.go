package models

import (
	"context"
	"github.com/uptrace/bun"
	"time"
)

// OauthClient contains the oauth clients
type OauthClient struct {
	ID          int64        `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt   time.Time    `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt   time.Time    `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	Secret      []byte       `validate:"-" bun:",nullzero,notnull"`
	Domain      string       `validate:"required,url" bun:",nullzero,notnull"`
	Description string       `validate:"-" bun:",nullzero"`
	OwnerID     int64        `validate:"-" bun:",nullzero,notnull"`
	Owner       *FediAccount `validate:"-" bun:"rel:belongs-to"`
}

var _ bun.BeforeAppendModelHook = (*OauthClient)(nil)

// BeforeAppendModel runs before a bun append operation
func (f *OauthClient) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		now := time.Now()
		f.CreatedAt = now
		f.UpdatedAt = now

		err := validate.Struct(f)
		if err != nil {
			return err
		}
	case *bun.UpdateQuery:
		f.UpdatedAt = time.Now()

		err := validate.Struct(f)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetSecret returns unencrypted secret
func (f *OauthClient) GetSecret() (string, error) {
	data, err := decrypt(f.Secret)
	return string(data), err
}

// SetSecret sets encrypted secret
func (f *OauthClient) SetSecret(s string) error {
	data, err := encrypt([]byte(s))
	if err != nil {
		return err
	}
	f.Secret = data
	return nil
}
