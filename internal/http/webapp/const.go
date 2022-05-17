package webapp

const (
	defaultCount = 10

	// DirStatic is the location of static assets.
	DirStatic = "static"

	// FormClientID is the key for a client-id form field.
	FormClientID = "client-id"
	// FormDescription is the key for a description form field.
	FormDescription = "description"
	// FormRedirectURI is the key for a redirect uri form field.
	FormRedirectURI = "redirect-uri"
	// FormToken is the key for a token form field.
	FormToken = "token"

	// XOriginAnonymous is a an cross-origin value for anonymous.
	XOriginAnonymous = "anonymous"
)

// SessionKey is a key used for storing data in a web session.
type SessionKey int

const (
	// SessionKeyAccountID contains the id of the currently logged-in user.
	SessionKeyAccountID SessionKey = iota
	// SessionKeyLoginRedirect contains the url to be redirected too after logging in.
	SessionKeyLoginRedirect
	// SessionKeyReturnURI contains the url to be redirected too after logging in.
	SessionKeyReturnURI
)
