package webapp

import (
	"github.com/feditools/go-lib/language"
	libtemplate "github.com/feditools/go-lib/template"
	"github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/path"
	"github.com/feditools/login/internal/template"
	nethttp "net/http"
)

// AdminOauthGetHandler serves the admin oauth page
func (m *Module) AdminOauthGetHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	l := logger.WithField("func", "AdminOauthGetHandler")

	// get localizer
	localizer := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer)

	// Init template variables
	tmplVars := &template.AdminOauth{
		Common: template.Common{
			PageTitle: localizer.TextOauth().String(),
		},
		Admin: template.Admin{
			Sidebar: makeAdminOauthSidebar(r),
		},
	}

	err := m.initTemplateAdmin(w, r, tmplVars)
	if err != nil {
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
		return
	}

	err = m.executeTemplate(w, template.AdminOauthName, tmplVars)
	if err != nil {
		l.Errorf("could not render %s template: %s", template.AdminOauthName, err.Error())
	}
}

func makeAdminOauthSidebar(r *nethttp.Request) libtemplate.Sidebar {
	// get localizer
	localizer := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer)

	// create sidebar
	newSidebar := libtemplate.Sidebar{
		{
			Text: localizer.TextOauth20Settings().String(),
			Children: []libtemplate.SidebarNode{
				{
					Text:    localizer.TextClient(2).String(),
					Matcher: path.ReAdminOauthClientsPre,
					Icon:    "desktop",
					URI:     path.AdminOauthClients,
				},
			},
		},
	}

	newSidebar.ActivateFromPath(r.URL.Path)

	return newSidebar
}
