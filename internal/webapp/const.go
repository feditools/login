package webapp

const (
	defaultCount = 10

	// DirStatic is the location of static assets
	DirStatic = DirWeb + "/static"
	// DirWeb is the location of files
	DirWeb = "web"

	// FormDescription is the ket for a description form field
	FormDescription = "description"
	// FormRedirectURI is the ket for a redirect uri form field
	FormRedirectURI = "redirect-uri"
)

// SessionKey is a key used for storing data in a web session
type SessionKey int

const (
	// SessionKeyAccountID contains the id of the currently logged-in user
	SessionKeyAccountID SessionKey = iota
	// SessionKeyLoginRedirect contains the url to be redirected too after logging in
	SessionKeyLoginRedirect
)
