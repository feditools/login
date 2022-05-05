package models

import (
	"time"
)

// OauthClient contains the oauth clients
type OauthClient struct {
	ID          int64        `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt   time.Time    `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt   time.Time    `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	Secret      []byte       `validate:"-" bun:",nullzero,notnull"`
	RedirectURI string       `validate:"required,url" bun:",nullzero,notnull"`
	Description string       `validate:"-" bun:",nullzero,notnull"`
	OwnerID     int64        `validate:"-" bun:",nullzero,notnull"`
	Owner       *FediAccount `validate:"-" bun:"rel:belongs-to"`
}
