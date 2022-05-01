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

	// PartAdd is used in a path for adding data
	PartAdd = "add"
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
	// PartLogout is used in a path for logout
	PartLogout = "logout"
	// PartMe is used in a path for the self
	PartMe = "me"
	// PartOauth is used in a path for oauth
	PartOauth = "oauth"
	// PartStatic is used in a path for static files
	PartStatic = "static"
	// PartToken is used in a path for static files
	PartToken = "token"

	// paths

	// Admin is the path for the admin page
	Admin = "/" + PartAdmin
	// AdminOauth is the sub path for the oauth admin page
	AdminOauth = Admin + AdminSubOauth
	// AdminOauthClients is the sub path for the oauth clients admin page
	AdminOauthClients = Admin + AdminSubOauthClients
	// AdminOauthClientAdd is the sub path for the oauth clients add admin page
	AdminOauthClientAdd = Admin + "/" + AdminSubOauthClientAdd
	// AdminSubOauth is the sub path for the oauth admin page
	AdminSubOauth = "/" + PartOauth
	// AdminSubOauthClient is the sub path for the oauth clients view client admin page
	AdminSubOauthClient = AdminSubOauthClients + "/" + VarClient
	// AdminSubOauthClients is the sub path for the oauth clients admin page
	AdminSubOauthClients = AdminSubOauth + "/" + PartClients
	// AdminSubOauthClientAdd is the sub path for the oauth clients add admin page
	AdminSubOauthClientAdd = AdminSubOauthClients + "/" + PartAdd
	// CallbackOauth is the path for an oauth callback
	CallbackOauth = "/" + PartCallback + "/" + PartOauth + "/" + VarInstance
	// Login is the path for the login page
	Login = "/" + PartLogin
	// Logout is the path for the logout page
	Logout = "/" + PartLogout
	// Me is the path for getting the logged-in user
	Me = "/" + PartMe
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

	// VarClientID is the id of the client variable
	VarClientID = "instance"
	// VarClient is the var path of the client variable
	VarClient = "{" + VarClientID + ":" + reToken + "}"
	// VarInstanceID is the id of the instance variable
	VarInstanceID = "instance"
	// VarInstance is the var path of the instance variable
	VarInstance = "{" + VarInstanceID + ":" + reToken + "}"
)
