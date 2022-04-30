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

	// PartAdmin is used in a path for administrative tasks
	PartAdmin = "admin"
	// PartAuthorize is used in a path for authorization
	PartAuthorize = "authorize"
	// PartCallback is used in a path for callback
	PartCallback = "callback"
	// PartClients is used in a path for oauth clients
	PartClients = "clients"
	// PartLogin is used in a path for login
	PartLogin = "login"
	// PartOauth is used in a path for oauth
	PartOauth = "oauth"
	// PartStatic is used in a path for static files
	PartStatic = "static"
	// PartToken is used in a path for static files
	PartToken = "token"

	// paths

	// Admin is the path for the admin page
	Admin = "/" + PartAdmin
	// AdminOauthClients is the sub path for the oauth clients admin page
	AdminOauthClients = Admin + AdminSubOauthClients
	// AdminSubOauthClients is the sub path for the oauth clients admin page
	AdminSubOauthClients = "/" + PartClients
	// CallbackOauth is the path for an oauth callback
	CallbackOauth = "/" + PartCallback + "/" + PartOauth + "/" + VarInstance
	// Login is the path for the login page
	Login = "/" + PartLogin
	// Oauth is the path prefix for oauth
	Oauth = "/" + PartOauth
	// OauthAuthorize is the path prefix for the oauth authorization
	OauthAuthorize = Oauth + "/" + PartAuthorize
	// OauthToken is the path prefix for the oauth token
	OauthToken = Oauth + "/" + PartToken
	// Static is the path for static files
	Static = "/" + PartStatic + "/"
	// StaticCSS is the path
	StaticCSS = Static + "css"

	// vars

	// VarInstanceID is the id of the instance variable
	VarInstanceID = "instance"
	// VarInstance is the var path of the instance variable
	VarInstance = "{" + VarInstanceID + ":" + reToken + "}"
)
