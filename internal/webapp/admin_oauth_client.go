package webapp

import (
	"github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/language"
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
	}

	// make admin navbar
	localizer := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer)
	navbar := makeAdminNavbar(r, localizer)
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
