package models

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

// OauthScope contains the oauth scopes.
type OauthScope struct {
	ID          int64     `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt   time.Time `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt   time.Time `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	Name        string    `validate:"required" bun:",nullzero,notnull"`
	Description string    `validate:"-" bun:",nullzero"`
	Default     bool      `validate:"-" bun:",notnull,default:false"`
}

var _ bun.BeforeAppendModelHook = (*OauthScope)(nil)

// BeforeAppendModel runs before a bun append operation.
func (f *OauthScope) BeforeAppendModel(_ context.Context, query bun.Query) error {
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
