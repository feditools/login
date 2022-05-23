package kv

import "testing"

func TestKeyFediActor(t *testing.T) {
	want := "login:fedi:actor:https://example.com/actor"
	v := KeyFediActor("https://example.com/actor")
	if v != want {
		t.Errorf("enexpected value for KeyFediActor, got: '%s', want: '%s'.", v, want)
	}
}

func TestKeyFediNodeInfo(t *testing.T) {
	want := "login:fedi:ni:example.com"
	v := KeyFediNodeInfo("example.com")
	if v != want {
		t.Errorf("enexpected value for TestKeyFediNodeInfo, got: '%s', want: '%s'.", v, want)
	}
}

func TestKeyNonceToken(t *testing.T) {
	want := "login:oauth:nonce:42:testtest1234"
	v := KeyOauthNonce("42", "testtest1234")
	if v != want {
		t.Errorf("enexpected value for TestKeyNonceToken, got: '%s', want: '%s'.", v, want)
	}
}

func TestKeyOauthToken(t *testing.T) {
	want := "login:oauth:token:"
	v := KeyOauthToken()
	if v != want {
		t.Errorf("enexpected value for KeyOauthToken, got: '%s', want: '%s'.", v, want)
	}
}

func TestKeySession(t *testing.T) {
	want := "login:session:" //nolint
	v := KeySession()
	if v != want {
		t.Errorf("enexpected value for KeyDomains, got: '%s', want: '%s'.", v, want)
	}
}
