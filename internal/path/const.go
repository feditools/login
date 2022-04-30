package path

const (
	// files

	// FileDefaultCSS is the css document applies to all pages
	FileDefaultCSS = StaticCSS + "/default.min.css"
	// FileErrorCSS is the css document applies to the error page
	FileErrorCSS = StaticCSS + "/error.min.css"
	// FileLoginCSS is the css document applies to the login page
	FileLoginCSS = StaticCSS + "/login.min.css"

	// parts

	// PartCallback is used in a path for callback
	PartCallback = "callback"
	// PartLogin is used in a path for login
	PartLogin = "login"
	// PartOauth is used in a path for oauth
	PartOauth = "oauth"
	// PartStatic is used in a path for static files
	PartStatic = "static"

	// paths

	// CallbackOauth is the path for an oauth callback
	CallbackOauth = "/" + PartCallback + "/" + PartOauth + "/" + VarInstance
	// Login is the path for the login page
	Login = "/" + PartLogin
	// Static is the path for static files
	Static = "/" + PartStatic + "/"
	// StaticCSS is the path
	StaticCSS = Static + "css"

	// regexes

	// ReToken is regex to match a token
	ReToken = `[a-zA-Z0-9_]{16,}`

	// vars

	// VarInstanceID is the id of the instance variable
	VarInstanceID = "instance"
	// VarInstance is the var path of the instance variable
	VarInstance = "{" + VarInstanceID + ":" + ReToken + "}"
)
