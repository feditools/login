package models

import "time"

// FediAccount represents a federated social account
type FediAccount struct {
	ID                   int64         `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt            time.Time     `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt            time.Time     `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	Username             string        `validate:"-" bun:",unique:unique_fedi_user,nullzero,notnull"`
	InstanceID           int64         `validate:"-" bun:",unique:unique_fedi_user,nullzero,notnull"`
	Instance             *FediInstance `validate:"-" bun:"rel:belongs-to,join:instance_id=id"`
	DisplayName          string        `validate:"-" bun:",notnull"`
	DisplayNameUpdatedAt time.Time     `validate:"-" bun:",notnull"`
	SignInCount          int           `validate:"min=0" bun:",notnull,default:0"`
	AccessToken          []byte        `validate:"-" bun:",nullzero"`

	// login stuff
	Admin bool `validate:"-" bun:",notnull,default:false"`
}
