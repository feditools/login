package kv

const (
	keyBase = "login:"

	keyFedi         = keyBase + "fedi:"
	keyFediActor    = keyFedi + "actor:"
	keyFediNodeInfo = keyFedi + "ni:"

	keyOauth      = keyBase + "oauth:"
	keyOauthNonce = keyOauth + "nonce:"
	keyOauthToken = keyOauth + "token:"

	keySession = keyBase + "session:"
)

// KeyFediActor returns the kv key which holds cached actor.
func KeyFediActor(u string) string { return keyFediActor + u }

// KeyFediNodeInfo returns the kv key which holds cached nodeinfo.
func KeyFediNodeInfo(d string) string { return keyFediNodeInfo + d }

// KeyOauthNonce returns the kv key which holds oauth nonce.
func KeyOauthNonce(uid string, sid string) string { return keyOauthNonce + uid + ":" + sid }

// KeyOauthToken returns the oauth token key prefix.
func KeyOauthToken() string { return keyOauthToken }

// KeySession returns the base kv key prefix.
func KeySession() string { return keySession }
