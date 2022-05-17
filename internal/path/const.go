package path

const (
	// files.

	// FileDefaultCSS is the css document applies to all pages.
	FileDefaultCSS = StaticCSS + "/default.min.css"
	// FileErrorCSS is the css document applies to the error page.
	FileErrorCSS = StaticCSS + "/error.min.css"
	// FileLoginCSS is the css document applies to the login page.
	FileLoginCSS = StaticCSS + "/login.min.css"

	// parts.

	// PartAccounts is used in a path for accounts.
	PartAccounts = "accounts"
	// PartAdd is used in a path for adding data.
	PartAdd = "add"
	// PartAdmin is used in a path for administrative tasks.
	PartAdmin = "admin"
	// PartApplicationTokens is used in a path for application tokens.
	PartApplicationTokens = "app-tokens"
	// PartAuthorize is used in a path for authorization.
	PartAuthorize = "authorize"
	// PartCallback is used in a path for callback.
	PartCallback = "callback"
	// PartClients is used in a path for oauth clients.
	PartClients = "clients"
	// PartFediverse is used in a path for federated things.
	PartFediverse = "fedi"
	// PartInstances is used in a path for instances.
	PartInstances = "instances"
	// PartLogin is used in a path for login.
	PartLogin = "login"
	// PartLogout is used in a path for logout.
	PartLogout = "logout"
	// PartMe is used in a path for the self.
	PartMe = "me"
	// PartOauth is used in a path for oauth.
	PartOauth = "oauth"
	// PartStatic is used in a path for static files.
	PartStatic = "static"
	// PartSystem is used in a path for system things.
	PartSystem = "system"
	// PartToken is used in a path for static files.
	PartToken = "token"

	// paths.

	// Admin is the path for the admin page.
	Admin = "/" + PartAdmin

	// AdminFediverse is the path for the fediverse admin page.
	AdminFediverse = Admin + AdminSubFediverse
	// AdminFediverseAccounts is the path for the fediverse admin page.
	AdminFediverseAccounts = Admin + AdminSubFediverseAccounts
	// AdminFediverseInstances is the path for the fediverse instances page.
	AdminFediverseInstances = Admin + AdminSubFediverseInstances

	// AdminSubFediverse is the sub path for the fediverse admin page.
	AdminSubFediverse = "/" + PartFediverse
	// AdminSubFediverseAccounts is the sub path for the fediverse admin accounts page.
	AdminSubFediverseAccounts = AdminSubFediverse + "/" + PartAccounts
	// AdminSubFediverseInstances is the sub path for the fediverse admin instances page.
	AdminSubFediverseInstances = AdminSubFediverse + "/" + PartInstances

	// AdminOauth is the path for the oauth admin page.
	AdminOauth = Admin + AdminSubOauth
	// AdminOauthClients is the path for the oauth clients admin page.
	AdminOauthClients = Admin + AdminSubOauthClients
	// AdminOauthClientAdd is the path for the oauth clients add admin page.
	AdminOauthClientAdd = Admin + AdminSubOauthClientAdd

	// AdminSubOauth is the sub path for the oauth admin page.
	AdminSubOauth = "/" + PartOauth
	// AdminSubOauthClient is the sub path for the oauth clients view client admin page.
	AdminSubOauthClient = AdminSubOauthClients + "/" + VarClient
	// AdminSubOauthClients is the sub path for the oauth clients admin page.
	AdminSubOauthClients = AdminSubOauth + "/" + PartClients
	// AdminSubOauthClientAdd is the sub path for the oauth clients add admin page.
	AdminSubOauthClientAdd = AdminSubOauthClients + "/" + PartAdd

	// AdminSystem is the path for the system admin page.
	AdminSystem = Admin + AdminSubSystem
	// AdminSystemApplicationTokens is the path for the application tokens page.
	AdminSystemApplicationTokens = Admin + AdminSubSystemApplicationTokens
	// AdminSystemApplicationTokenAdd is the path for the add application tokens page.
	AdminSystemApplicationTokenAdd = Admin + AdminSubSystemApplicationTokenAdd

	// AdminSubSystem is the sub path for the system admin page.
	AdminSubSystem = "/" + PartSystem
	// AdminSubSystemApplicationTokens is the sub path for the application tokens page.
	AdminSubSystemApplicationTokens = AdminSubSystem + "/" + PartApplicationTokens
	// AdminSubSystemApplicationTokenAdd is the sub path for the add application tokens page.
	AdminSubSystemApplicationTokenAdd = AdminSubSystemApplicationTokens + "/" + PartAdd

	// CallbackOauth is the path for an oauth callback.
	CallbackOauth = "/" + PartCallback + "/" + PartOauth + "/" + VarInstance
	// Login is the path for the login page.
	Login = "/" + PartLogin
	// Logout is the path for the logout page.
	Logout = "/" + PartLogout
	// Me is the path for getting the logged-in user.
	Me = "/" + PartMe
	// Oauth is the path prefix for oauth.
	Oauth = "/" + PartOauth
	// OauthAuthorize is the path prefix for the oauth authorization.
	OauthAuthorize = Oauth + "/" + PartAuthorize
	// OauthToken is the path prefix for the oauth token.
	OauthToken = Oauth + "/" + PartToken
	// Static is the path for static files.
	Static = "/" + PartStatic + "/"
	// StaticCSS is the path.
	StaticCSS = Static + "css"

	// vars.

	// VarClientID is the id of the client variable.
	VarClientID = "instance"
	// VarClient is the var path of the client variable.
	VarClient = "{" + VarClientID + ":" + reToken + "}"
	// VarInstanceID is the id of the instance variable.
	VarInstanceID = "instance"
	// VarInstance is the var path of the instance variable.
	VarInstance = "{" + VarInstanceID + ":" + reToken + "}"
)
