package template

// AdminOauthClientName is the name of the admin oauth clients template
const AdminOauthClientName = "admin_oauth_client"

// AdminOauthClient contains the variables for the "admin_oauth_client" template.
type AdminOauthClient struct {
	Common
	Admin

	HrefAddClient string
}
