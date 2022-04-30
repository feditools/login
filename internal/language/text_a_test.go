package language

import (
	"fmt"
	"golang.org/x/text/language"
	"testing"
)

func TestTextAccount(t *testing.T) {
	langMod, _ := New()

	tables := []struct {
		x language.Tag
		c int
		n string
		l language.Tag
	}{
		{language.English, 1, "Account", language.English},
		{language.English, 2, "Accounts", language.English},
		{language.Spanish, 1, "Account", language.English},
		{language.Spanish, 2, "Accounts", language.English},
		{language.Hindi, 1, "Account", language.English},
		{language.Hindi, 2, "Accounts", language.English},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Translating to %s", i, table.x)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			localizer, err := langMod.NewLocalizer(table.x.String())
			if err != nil {
				t.Errorf("[%d] can't get localizer for %s: %s", i, table.x, err.Error())
				return
			}

			result := localizer.TextAccount(table.c)
			if result.String() != table.n {
				t.Errorf("[%d] got invalid translation for %s, got: %v, want: %v,", i, table.x, result.String(), table.n)
			}
			if result.Language() != table.l {
				t.Errorf("[%d] got invalid language for %s, got: %v, want: %v,", i, table.x, result.Language(), table.l)
			}
		})
	}
}
