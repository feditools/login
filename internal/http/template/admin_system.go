package template

import (
	libtemplate "github.com/feditools/go-lib/template"
	"github.com/feditools/login/internal/models"
)

// AdminSystemName is the name of the admin oauth template.
const AdminSystemName = "admin_system"

// AdminSystem contains the variables for the admin oauth template.
type AdminSystem struct {
	Common
	Admin
}

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

// AdminSystemApplicationTokenAddName is the name of the admin system application add token template.
const AdminSystemApplicationTokenAddName = "admin_system_applicationtoken_add"

// AdminSystemApplicationTokenAdd contains the variables for the admin system application add token template.
type AdminSystemApplicationTokenAdd struct {
	Common
	Admin

	FormInputDescription *libtemplate.FormInput
	FormButtonSubmitText string
}

// AdminSystemApplicationTokenAddedName is the name of the admin system application token added template.
const AdminSystemApplicationTokenAddedName = "admin_system_applicationtoken_added"

// AdminSystemApplicationTokenAdded contains the variables for the admin system application token added template.
type AdminSystemApplicationTokenAdded struct {
	Common
	Admin

	FormInputDescription *libtemplate.FormInput
	FormInputToken       *libtemplate.FormInput
}
