package models

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

// ApplicationToken contains application authentication tokens.
type ApplicationToken struct {
	ID          int64        `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt   time.Time    `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt   time.Time    `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	Token       string       `validate:"-" bun:",nullzero,notnull"`
	Description string       `validate:"-" bun:",nullzero,notnull"`
	CreatedByID int64        `validate:"-" bun:",nullzero,notnull"`
	CreatedBy   *FediAccount `validate:"-" bun:"rel:belongs-to"`
}

var _ bun.BeforeAppendModelHook = (*ApplicationToken)(nil)

// BeforeAppendModel runs before a bun append operation.
func (f *ApplicationToken) BeforeAppendModel(_ context.Context, query bun.Query) error {
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
