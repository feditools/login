package kv

const (
	keyBase = "login:"

	keyFedi         = keyBase + "fedi:"
	keyFediActor    = keyFedi + "actor:"
	keyFediNodeInfo = keyFedi + "ni:"

	keyOauth             = keyBase + "oauth:"
	keyOauthNonce        = keyOauth + "nonce:"
	keyOauthNonceLogin   = keyOauthNonce + "login:"
	keyOauthNonceRefresh = keyOauthNonce + "refresh:"
	keyOauthToken        = keyOauth + "token:"

	keySession = keyBase + "session:"
)

// KeyFediActor returns the kv key which holds cached actor.
func KeyFediActor(u string) string { return keyFediActor + u }

// KeyFediNodeInfo returns the kv key which holds cached nodeinfo.
func KeyFediNodeInfo(d string) string { return keyFediNodeInfo + d }

// KeyOauthNonceLogin returns the kv key which holds oauth nonce received from a login request.
func KeyOauthNonceLogin(uid string) string { return keyOauthNonceLogin + uid }

// KeyOauthNonceRefresh returns the kv key which holds oauth nonce tied to a refresh token.
func KeyOauthNonceRefresh(refreshToken string) string { return keyOauthNonceRefresh + refreshToken }

// KeyOauthToken returns the oauth token key prefix.
func KeyOauthToken() string { return keyOauthToken }

// KeySession returns the base kv key prefix.
func KeySession() string { return keySession }
