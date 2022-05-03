package template

// AdminOauthClientName is the name of the admin oauth clients template
const AdminOauthClientName = "admin_oauth_client"

// AdminOauthClient contains the variables for the "admin_oauth_client" template.
type AdminOauthClient struct {
	Common
	Admin

	HrefAddClient string
}

// AdminOauthClientAddName is the name of the admin oauth clients template
const AdminOauthClientAddName = "admin_oauth_client_add"

// AdminOauthClientAdd contains the variables for the "admin_oauth_client" template.
type AdminOauthClientAdd struct {
	Common
	Admin

	FormInputDescription *FormInput
	FormInputReturnURI   *FormInput
	FormButtonSubmitText string
}

// AdminOauthClientAddedName is the name of the admin oauth added clients template
const AdminOauthClientAddedName = "admin_oauth_client_added"

// AdminOauthClientAdded contains the variables for the "admin_oauth_client" template.
type AdminOauthClientAdded struct {
	Common
	Admin

	FormInputDescription  *FormInput
	FormInputClientID     *FormInput
	FormInputClientSecret *FormInput
	FormInputReturnURI    *FormInput
}
