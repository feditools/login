package webapp

import (
	"github.com/feditools/go-lib/language"
	libtemplate "github.com/feditools/go-lib/template"
	"github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/path"
	nethttp "net/http"
)

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
