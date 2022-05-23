package models

import (
	"context"
	"testing"

	"github.com/uptrace/bun"
)

func TestFediInstance_BeforeAppendModel_Insert(t *testing.T) {
	obj := &FediInstance{
		Domain:   "test.com",
		ActorURI: "https://test.com/actor",
	}

	err := obj.BeforeAppendModel(context.Background(), &bun.InsertQuery{})
	if err != nil {
		t.Errorf("got error: %s", err.Error())

		return
	}

	if obj.CreatedAt.Equal(testEmptyTime) {
		t.Errorf("invalid created at time: %s", obj.CreatedAt.String())
	}
	if obj.UpdatedAt.Equal(testEmptyTime) {
		t.Errorf("invalid updated at time: %s", obj.UpdatedAt.String())
	}
}

func TestFediInstance_BeforeAppendModel_Update(t *testing.T) {
	obj := &FediInstance{
		Domain:   "test.com",
		ActorURI: "https://test.com/actor",
	}

	err := obj.BeforeAppendModel(context.Background(), &bun.UpdateQuery{})
	if err != nil {
		t.Errorf("got error: %s", err.Error())

		return
	}

	if obj.UpdatedAt.Equal(testEmptyTime) {
		t.Errorf("invalid updated at time: %s", obj.UpdatedAt.String())
	}
}
