package http

// SessionKey is a key used for storing data in a web session.
type SessionKey int

const (
	// SessionKeyAccountID contains the id of the currently logged-in user.
	SessionKeyAccountID SessionKey = iota
	// SessionKeyLoginRedirect contains the url to be redirected too after logging in.
	SessionKeyLoginRedirect
	// SessionKeyReturnURI contains the url to be redirected too after logging in.
	SessionKeyReturnURI
	// SessionKeyOAuthNonce contains the nonce sent to the oauth server.
	SessionKeyOAuthNonce
)
