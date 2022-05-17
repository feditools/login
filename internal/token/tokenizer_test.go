package token

import (
	"errors"
	"fmt"
	"testing"

	"github.com/feditools/login/internal/config"
	"github.com/feditools/login/internal/models"
	"github.com/spf13/viper"
)

//revive:disable:add-constant

var testTables = []struct {
	k Kind
	i int64
	t string
}{
	{KindApplicationToken, 22, "42L81QrOF0aorZlq"},
	{KindFediAccount, 69420, "n4jbQJEimvjP5WZy"},
	{KindFediInstance, 84685, "MeLrPDrcOg5agV8K"},
	{KindOauthClient, 184, "p6jKkP6VIlgPqrEl"},
	{KindOauthScope, 1, "pwM9xPkwuZP7ve0n"},
}

func TestNew(t *testing.T) {
	viper.Reset()

	viper.Set(config.Keys.TokenSalt, "test1234")

	tokenizer, err := New()
	if err != nil {
		t.Errorf("got error: %s", err.Error())
		return
	}

	if tokenizer.h == nil {
		t.Errorf("hashid is nil")
		return
	}
}

func TestNew_SaltEmpty(t *testing.T) {
	viper.Reset()

	tokenizer, err := New()
	if !errors.Is(err, ErrSaltEmpty) {
		t.Errorf("unexpected error, got: '%s', want: '%s'", err, ErrSaltEmpty)
		return
	}

	if tokenizer != nil {
		t.Errorf("unexpected tokenizer, got: '%T', want: '%T'", tokenizer, nil)
		return
	}
}

func TestTokenizer_DecodeToken(t *testing.T) {
	tokenizer, err := testNewTestTokenizer()
	if err != nil {
		t.Errorf("init: %s", err.Error())
		return
	}

	for i, table := range testTables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Running EncodeToken %d(%s)", i, table.i, table.k.String())
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			kind, id, err := tokenizer.DecodeToken(table.t)
			if err != nil {
				t.Errorf("got error: %s", err.Error())
				return
			}
			if kind != table.k {
				t.Errorf("[%d] wrong kind: got '%s', want '%s'", i, kind, table.k)
			}
			if id != table.i {
				t.Errorf("[%d] wrong id: got '%d', want '%d'", i, id, table.i)
			}
		})
	}
}

func TestTokenizer_DecodeToken_InvalidLength(t *testing.T) {
	tokenizer, err := testNewTestTokenizer()
	if err != nil {
		t.Errorf("init: %s", err.Error())
		return
	}

	_, _, err = tokenizer.DecodeToken("1vxqadgcYibQ2pOj")
	errText := "negative number not supported"
	if !errors.Is(err, ErrInvalidLength) {
		t.Errorf("unexpected error, got: '%s', want: '%s'", err, errText)
		return
	}
}

func TestTokenizer_EncodeToken(t *testing.T) {
	tokenizer, err := testNewTestTokenizer()
	if err != nil {
		t.Errorf("init: %s", err.Error())
		return
	}

	for i, table := range testTables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Running EncodeToken %d(%s)", i, table.i, table.k.String())
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			token, err := tokenizer.EncodeToken(table.k, table.i)
			if err != nil {
				t.Errorf("got error: %s", err.Error())
				return
			}
			if token != table.t {
				t.Errorf("[%d] wrong token: got '%s', want '%s'", i, token, table.t)
			}
		})
	}
}

func TestTokenizer_EncodeToken_Negative(t *testing.T) {
	tokenizer, err := testNewTestTokenizer()
	if err != nil {
		t.Errorf("init: %s", err.Error())
		return
	}

	_, err = tokenizer.EncodeToken(KindFediAccount, -1)
	errText := "negative number not supported"
	if err == nil {
		t.Errorf("expected error, got: 'nil', want: '%s'", errText)
		return
	}
	if err.Error() != errText {
		t.Errorf("unexpected error, got: '%s', want: '%s'", err, errText)
		return
	}
}

func TestTokenizer_GetToken(t *testing.T) {
	tokenizer, err := testNewTestTokenizer()
	if err != nil {
		t.Errorf("init: %s", err.Error())
		return
	}

	tables := []struct {
		o interface{}
		t string
	}{
		{models.FediAccount{ID: 1}, "RNrm2XxAiGPGpyD4"},
		{&models.FediAccount{ID: 1}, "RNrm2XxAiGPGpyD4"},
		{models.FediInstance{ID: 1}, "5yp6YXmmcLX9Kgro"},
		{&models.FediInstance{ID: 1}, "5yp6YXmmcLX9Kgro"},
		{models.OauthClient{ID: 1}, "pMeLrPDzIxagV8KY"},
		{&models.OauthClient{ID: 1}, "pMeLrPDzIxagV8KY"},
		{models.OauthScope{ID: 1}, "pwM9xPkwuZP7ve0n"},
		{&models.OauthScope{ID: 1}, "pwM9xPkwuZP7ve0n"},
		{&struct{}{}, ""},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Running GetToken for %T", i, table.o)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			token := tokenizer.GetToken(table.o)
			if token != table.t {
				t.Errorf("[%d] wrong token: got '%s', want '%s'", i, token, table.t)
			}
		})
	}
}

func testNewTestTokenizer() (*Tokenizer, error) {
	viper.Reset()
	viper.Set(config.Keys.TokenSalt, "test1234")
	return New()
}

func BenchmarkTokenizer_DecodeToken(b *testing.B) {
	tokenizer, err := testNewTestTokenizer()
	if err != nil {
		b.Errorf("init: %s", err.Error())
		return
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _, _ = tokenizer.DecodeToken("AMroVP59cwLPE5pb")
		}
	})
}

func BenchmarkTokenizer_EncodeToken(b *testing.B) {
	tokenizer, err := testNewTestTokenizer()
	if err != nil {
		b.Errorf("init: %s", err.Error())
		return
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = tokenizer.EncodeToken(KindFediInstance, 123)
		}
	})
}

//revive:enable:add-constant
