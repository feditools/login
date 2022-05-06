package models

import "time"

// FediInstance represents a federated social instance.
type FediInstance struct {
	ID             int64     `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt      time.Time `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt      time.Time `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	Domain         string    `validate:"-" bun:",nullzero,notnull,unique"`
	ActorURI       string    `validate:"url" bun:",nullzero,notnull"`
	ServerHostname string    `validate:"-" bun:",nullzero,notnull,unique"`
	Software       string    `validate:"-" bun:",nullzero,notnull"`
	ClientID       string    `validate:"-" bun:",nullzero"`
	ClientSecret   []byte    `validate:"-" bun:",nullzero"`
}
