package template

import (
	libtemplate "github.com/feditools/go-lib/template"
	"github.com/feditools/login/internal/models"
)

// AdminOauthClientsName is the name of the admin oauth clients template
const AdminOauthClientsName = "admin_oauth_clients"

// AdminOauthClients contains the variables for the "admin_oauth_client" template.
type AdminOauthClients struct {
	Common
	Admin

	HrefAddClient string
	OauthClients  *[]*models.OauthClient
}

// AdminOauthClientAddName is the name of the admin oauth clients template
const AdminOauthClientAddName = "admin_oauth_client_add"

// AdminOauthClientAdd contains the variables for the "admin_oauth_client" template.
type AdminOauthClientAdd struct {
	Common
	Admin

	FormInputDescription *libtemplate.FormInput
	FormInputRedirectURI *libtemplate.FormInput
	FormButtonSubmitText string
}

// AdminOauthClientAddedName is the name of the admin oauth added clients template
const AdminOauthClientAddedName = "admin_oauth_client_added"

// AdminOauthClientAdded contains the variables for the "admin_oauth_client" template.
type AdminOauthClientAdded struct {
	Common
	Admin

	FormInputDescription  *libtemplate.FormInput
	FormInputClientID     *libtemplate.FormInput
	FormInputClientSecret *libtemplate.FormInput
	FormInputRedirectURI  *libtemplate.FormInput
}
