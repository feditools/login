package models

import (
	"time"
)

// OauthScope contains the oauth scopes
type OauthScope struct {
	ID          int64     `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt   time.Time `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt   time.Time `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	Name        string    `validate:"required" bun:",nullzero,notnull"`
	Description string    `validate:"-" bun:",nullzero"`
	Default     bool      `validate:"-" bun:",notnull,default:false"`
}
