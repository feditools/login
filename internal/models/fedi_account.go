package models

import (
	"context"
	"github.com/uptrace/bun"
	"time"
)

// FediAccount represents a federated social account.
type FediAccount struct {
	ID                   int64         `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt            time.Time     `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt            time.Time     `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	Username             string        `validate:"-" bun:",unique:unique_fedi_user,nullzero,notnull"`
	InstanceID           int64         `validate:"-" bun:",unique:unique_fedi_user,notnull,nullzero"`
	Instance             *FediInstance `validate:"-" bun:"rel:belongs-to,join:instance_id=id"`
	DisplayName          string        `validate:"-" bun:",notnull"`
	DisplayNameUpdatedAt time.Time     `validate:"-" bun:",notnull"`
	SignInCount          int           `validate:"min=0" bun:",notnull,default:0"`
	AccessToken          []byte        `validate:"-" bun:",nullzero"`
}

var _ bun.BeforeAppendModelHook = (*FediAccount)(nil)

// BeforeAppendModel runs before a bun append operation
func (f *FediAccount) BeforeAppendModel(_ context.Context, query bun.Query) error {
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

// GetAccessToken returns unencrypted access token
func (f *FediAccount) GetAccessToken() (string, error) {
	data, err := decrypt(f.AccessToken)
	return string(data), err
}

// SetAccessToken sets encrypted access token
func (f *FediAccount) SetAccessToken(a string) error {
	data, err := encrypt([]byte(a))
	if err != nil {
		return err
	}
	f.AccessToken = data
	return nil
}
