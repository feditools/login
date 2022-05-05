package template

import (
	libtemplate "github.com/feditools/go-lib/template"
	"github.com/feditools/login/internal/models"
)

// AdminFediInstancesName is the name of the admin oauth clients template
const AdminFediInstancesName = "admin_oauth_clients"

// AdminFediInstances contains the variables for the fedi instances template.
type AdminFediInstances struct {
	Common
	Admin

	HRefViewFediInstance string
	OauthClients         *[]*models.FediInstance
	Pagination           libtemplate.Pagination
}
