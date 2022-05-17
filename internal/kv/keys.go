package kv

const (
	keyBase = "login:"

	keyFedi         = keyBase + "fedi:"
	keyFediActor    = keyFedi + "actor:"
	keyFediNodeInfo = keyFedi + "ni:"

	keyOauth      = keyBase + "oauth:"
	keyOauthToken = keyOauth + "token:"

	keySession = keyBase + "session:"
)

// KeyFediActor returns the kv key which holds cached actor.
func KeyFediActor(u string) string { return keyFediActor + u }

// KeyFediNodeInfo returns the kv key which holds cached nodeinfo.
func KeyFediNodeInfo(d string) string { return keyFediNodeInfo + d }

// KeySession returns the base kv key prefix.
func KeySession() string { return keySession }

// KeyOauthToken returns the oauth token key prefix.
func KeyOauthToken() string { return keyOauthToken }
