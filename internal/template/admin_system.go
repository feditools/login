package template

import (
	libtemplate "github.com/feditools/go-lib/template"
	"github.com/feditools/login/internal/models"
)

// AdminSystemApplicationTokensName is the name of the admin system application tokens template.
const AdminSystemApplicationTokensName = "admin_system_applicationtokens"

// AdminSystemApplicationTokens contains the variables for the admin system application tokens template.
type AdminSystemApplicationTokens struct {
	Common
	Admin

	HRefAddApplicationToken  string
	HRefViewApplicationToken string
	HRefViewFediAccount      string
	ApplicationTokens        *[]*models.ApplicationToken
	Pagination               libtemplate.Pagination
}

// AdminApplicationTokenAddName is the name of the admin system application add token template
const AdminApplicationTokenAddName = "admin_system_applicationtoken_add"

// AdminApplicationTokenAdd contains the variables for the admin system application add token template.
type AdminApplicationTokenAdd struct {
	Common
	Admin

	FormInputDescription *libtemplate.FormInput
	FormButtonSubmitText string
}

// AdminApplicationTokenAddedName is the name of the admin system application token added template
const AdminApplicationTokenAddedName = "admin_system_applicationtoken_added"

// AdminApplicationTokenAdded contains the variables for the admin system application token added template.
type AdminApplicationTokenAdded struct {
	Common
	Admin

	FormInputDescription *libtemplate.FormInput
	FormInputToken       *libtemplate.FormInput
}
