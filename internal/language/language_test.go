package language

import "testing"

func TestNew(t *testing.T) {
	langMod, err := New()
	if err != nil {
		t.Errorf("can't get new language module: %s", err.Error())
		return
	}

	if langMod == nil {
		t.Errorf("language module is nil")
		return
	}

	if langMod.langBundle == nil {
		t.Errorf("language module's bundle is nil")
		return
	}

	if langMod.Language() != defaultLanguage {
		t.Errorf("got invalid language, got: %v, want: %v,", langMod.Language().String(), defaultLanguage.String())
		return
	}
}
