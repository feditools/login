package models

import "time"

// FediAccount represents a federated social account.
type FediAccount struct {
	ID          int64         `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt   time.Time     `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt   time.Time     `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	ActorURI    string        `validate:"url" bun:",nullzero,notnull"`
	Username    string        `validate:"-" bun:",unique:unique_fedi_user,nullzero,notnull"`
	InstanceID  int64         `validate:"-" bun:",unique:unique_fedi_user,nullzero,notnull"`
	Instance    *FediInstance `validate:"-" bun:"rel:belongs-to,join:instance_id=id"`
	DisplayName string        `validate:"-" bun:",notnull"`
	LastFinger  time.Time     `validate:"-" bun:",notnull"`
	LogInCount  int64         `validate:"-" bun:",nullzero,notnull,default:0"`
	LogInLast   time.Time     `validate:"-" bun:",nullzero"`
	AccessToken []byte        `validate:"-" bun:",nullzero"`

	// login stuff
	Admin bool `validate:"-" bun:",notnull"`
}
