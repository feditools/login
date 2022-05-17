package models

import "time"

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
