package kv

import "testing"

func TestKeyFediNodeInfo(t *testing.T) {
	want := "login:fedi:ni:example.com"
	v := KeyFediNodeInfo("example.com")
	if v != want {
		t.Errorf("enexpected value for KeyDomains, got: '%s', want: '%s'.", v, want)
	}
}

func TestKeySession(t *testing.T) {
	want := "login:session:"
	v := KeySession()
	if v != want {
		t.Errorf("enexpected value for KeyDomains, got: '%s', want: '%s'.", v, want)
	}
}
