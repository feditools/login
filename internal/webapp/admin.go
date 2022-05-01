package webapp

import (
	"github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/language"
	"github.com/feditools/login/internal/path"
	"github.com/feditools/login/internal/template"
	nethttp "net/http"
)

func makeAdminSidebar(r *nethttp.Request) template.Sidebar {
	// get localizer
	localizer := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer)

	// create sidebar
	newSidebar := template.Sidebar{
		{
			Children: []template.SidebarNode{
				{
					Text:     localizer.TextDashboard().String(),
					MatchStr: path.ReAdmin,
					Icon:     "home",
					URL:      path.Admin,
				},
			},
		},
		{
			Text: localizer.TextOauth20Settings().String(),
			Children: []template.SidebarNode{
				{
					Text:     localizer.TextClient(2).String(),
					MatchStr: path.ReAdminOauthClientsPre,
					Icon:     "desktop",
					URL:      path.AdminOauthClients,
				},
			},
		},
	}

	newSidebar.ActivateFromPath(r.URL.Path)

	return newSidebar
}

func makeAdminOauthSidebar(r *nethttp.Request) template.Sidebar {
	// get localizer
	localizer := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer)

	// create sidebar
	newSidebar := template.Sidebar{
		{
			Text: localizer.TextOauth20Settings().String(),
			Children: []template.SidebarNode{
				{
					Text:     localizer.TextClient(2).String(),
					MatchStr: path.ReAdminOauthClientsPre,
					Icon:     "desktop",
					URL:      path.AdminOauthClients,
				},
			},
		},
	}

	newSidebar.ActivateFromPath(r.URL.Path)

	return newSidebar
}
