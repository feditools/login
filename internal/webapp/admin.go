package webapp

import (
	"github.com/feditools/go-lib/language"
	libtemplate "github.com/feditools/go-lib/template"
	"github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/path"
	nethttp "net/http"
)

func makeAdminSidebar(r *nethttp.Request) libtemplate.Sidebar {
	// get localizer
	localizer := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer)

	// create sidebar
	newSidebar := libtemplate.Sidebar{
		{
			Children: []libtemplate.SidebarNode{
				{
					Text:    localizer.TextDashboard(1).String(),
					Matcher: path.ReAdmin,
					Icon:    "home",
					URI:     path.Admin,
				},
			},
		},
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
