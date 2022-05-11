package webapp

import (
	"github.com/feditools/go-lib/language"
	"github.com/feditools/go-lib/template"
	"github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/path"
	nethttp "net/http"
)

func makeAdminFediverseSidebar(r *nethttp.Request) libtemplate.Sidebar {
	// get localizer
	localizer := r.Context().Value(http.ContextKeyLocalizer).(*liblanguage.Localizer)

	// create sidebar
	newSidebar := libtemplate.Sidebar{
		{
			Text: localizer.TextOauth20Settings().String(),
			Children: []libtemplate.SidebarNode{
				{
					Text:    localizer.TextInstance(2).String(),
					Matcher: path.ReAdminFediverseInstancesPre,
					Icon:    "desktop",
					URI:     path.AdminFediverseInstances,
				},
				{
					Text:    localizer.TextAccount(2).String(),
					Matcher: path.ReAdminFediverseAccountsPre,
					Icon:    "user",
					URI:     path.AdminFediverseAccounts,
				},
			},
		},
	}

	newSidebar.ActivateFromPath(r.URL.Path)

	return newSidebar
}
