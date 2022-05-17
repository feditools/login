package template

import (
	libtemplate "github.com/feditools/go-lib/template"
	"github.com/feditools/login/internal/models"
)

// AdminFediName is the name of the admin fedi template.
const AdminFediName = "admin_fedi"

// AdminFedi contains the variables for the admin fedi template.
type AdminFedi struct {
	Common
	Admin
}

// AdminFediAccountsName is the name of the admin fedi accounts template.
const AdminFediAccountsName = "admin_fedi_accounts"

// AdminFediAccounts contains the variables for the fedi instances template.
type AdminFediAccounts struct {
	Common
	Admin

	HRefViewFediAccount  string
	HRefViewFediInstance string
	FediAccounts         *[]*models.FediAccount
	Pagination           libtemplate.Pagination
}

// AdminFediAccountsForInstanceName is the name of the admin fedi accounts for instance template.
const AdminFediAccountsForInstanceName = "admin_fedi_accounts_for_instance"

// AdminFediAccountsForInstance contains the variables for the fedi instances template.
type AdminFediAccountsForInstance struct {
	Common
	Admin

	HRefViewFediAccount  string
	HRefViewFediInstance string
	FediInstance         *models.FediInstance
	FediAccounts         *[]*models.FediAccount
	Pagination           libtemplate.Pagination
}

// AdminFediInstancesName is the name of the admin oauth clients template.
const AdminFediInstancesName = "admin_fedi_instances"

// AdminFediInstances contains the variables for the fedi instances template.
type AdminFediInstances struct {
	Common
	Admin

	HRefViewFediInstance      string
	FediInstances             *[]*models.FediInstance
	FediInstanceAccountCounts []int64
	Pagination                libtemplate.Pagination
}
