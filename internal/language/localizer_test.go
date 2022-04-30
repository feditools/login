package language

import "testing"

func TestNewLocalizer(t *testing.T) {
	langMod, _ := New()

	localizer, err := langMod.NewLocalizer()
	if err != nil {
		t.Errorf("can't get new language module: %s", err.Error())
		return
	}

	if localizer == nil {
		t.Errorf("localizer module is nil")
		return
	}

	if localizer.localizer == nil {
		t.Errorf("localizer module's localizer is nil")
		return
	}
}
