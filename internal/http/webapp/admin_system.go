package webapp

import (
	"github.com/feditools/go-lib/language"
	libtemplate "github.com/feditools/go-lib/template"
	"github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/http/template"
	"github.com/feditools/login/internal/path"
	nethttp "net/http"
)

// AdminSystemGetHandler serves the admin system page
func (m *Module) AdminSystemGetHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	l := logger.WithField("func", "AdminSystemGetHandler")

	// get localizer
	localizer := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer)

	// Init template variables
	tmplVars := &template.AdminSystem{
		Common: template.Common{
			PageTitle: localizer.TextSystem(1).String(),
		},
		Admin: template.Admin{
			Sidebar: makeAdminSystemSidebar(r),
		},
	}

	err := m.initTemplateAdmin(w, r, tmplVars)
	if err != nil {
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
		return
	}

	err = m.executeTemplate(w, template.AdminSystemName, tmplVars)
	if err != nil {
		l.Errorf("could not render %s template: %s", template.AdminSystemName, err.Error())
	}
}

func makeAdminSystemSidebar(r *nethttp.Request) libtemplate.Sidebar {
	// get localizer
	localizer := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer)

	// create sidebar
	newSidebar := libtemplate.Sidebar{
		{
			Children: []libtemplate.SidebarNode{
				{
					Text:    localizer.TextApplicationToken(2).String(),
					Matcher: path.ReAdminSystemApplicationTokensPre,
					Icon:    "desktop",
					URI:     path.AdminSystemApplicationTokens,
				},
			},
		},
	}

	newSidebar.ActivateFromPath(r.URL.Path)

	return newSidebar
}
