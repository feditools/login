package oauth

// SessionKey is a key used for storing data in a web session.
type SessionKey int

const (
	// SessionKeyState contains the state sent to the oauth server.
	SessionKeyState SessionKey = iota
	// SessionKeyCode contains the code sent to the oauth server.
	SessionKeyCode
	// SessionKeyNonce contains the nonce sent to the oauth server.
	SessionKeyNonce
)
