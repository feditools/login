package webapp

import (
	"github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/language"
	"github.com/feditools/login/internal/path"
	"github.com/feditools/login/internal/template"
	nethttp "net/http"
)

// AdminOauthClientsGetHandler serves the admin client page
func (m *Module) AdminOauthClientsGetHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	l := logger.WithField("func", "AdminOauthClientsGetHandler")

	// Init template variables
	tmplVars := &template.AdminOauthClient{
		Common: template.Common{
			PageTitle: "Admin Clients",
		},
		Admin: template.Admin{
			Sidebar: makeAdminOauthSidebar(r),
		},
		HrefAddClient: path.AdminOauthClientAdd,
	}

	// make admin navbar
	navbar := makeAdminNavbar(r)
	tmplVars.SetNavbar(navbar)

	err := m.initTemplate(w, r, tmplVars)
	if err != nil {
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
		return
	}

	err = m.executeTemplate(w, template.AdminOauthClientName, tmplVars)
	if err != nil {
		l.Errorf("could not render %s template: %s", template.AdminOauthClientName, err.Error())
	}
}

// AdminOauthClientAddGetHandler serves the admin add client page
func (m *Module) AdminOauthClientAddGetHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	l := logger.WithField("func", "AdminOauthClientAddGetHandler")

	// get localizer
	localizer := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer)

	// Init template variables
	tmplVars := &template.AdminOauthClientAdd{
		Common: template.Common{
			PageTitle: "Admin Add Clients",
		},
		Admin: template.Admin{
			Sidebar: makeAdminOauthSidebar(r),
		},

		FormInputDescriptionDisabled: false,
		FormInputDescriptionValue:    "",
		FormButtonSubmitText:         localizer.TextCreate().String(),
	}

	// make admin navbar
	navbar := makeAdminNavbar(r)
	tmplVars.SetNavbar(navbar)

	err := m.initTemplate(w, r, tmplVars)
	if err != nil {
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
		return
	}

	err = m.executeTemplate(w, template.AdminOauthClientAddName, tmplVars)
	if err != nil {
		l.Errorf("could not render %s template: %s", template.AdminOauthClientAddName, err.Error())
	}
}
