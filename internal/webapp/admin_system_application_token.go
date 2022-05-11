package webapp

import (
	"github.com/feditools/go-lib"
	"github.com/feditools/go-lib/language"
	libtemplate "github.com/feditools/go-lib/template"
	"github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/path"
	"github.com/feditools/login/internal/template"
	nethttp "net/http"
)

// AdminSystemApplicationTokensGetHandler serves the admin client page
func (m *Module) AdminSystemApplicationTokensGetHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	l := logger.WithField("func", "AdminSystemApplicationTokensGetHandler")

	// get localizer
	localizer := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer)

	// Init template variables
	tmplVars := &template.AdminSystemApplicationTokens{
		Common: template.Common{
			PageTitle: localizer.TextApplicationToken(2).String(),
		},
		Admin: template.Admin{
			Sidebar: makeAdminSystemSidebar(r),
		},
		HRefAddApplicationToken:  path.AdminSystemApplicationTokenAdd,
		HRefViewApplicationToken: path.AdminSystemApplicationTokens,
		HRefViewFediAccount:      path.AdminFediverseAccounts,
	}

	err := m.initTemplateAdmin(w, r, tmplVars)
	if err != nil {
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
		return
	}

	page, count, countFound := lib.GetPaginationFromURL(r.URL, defaultCount)

	// get oauth clients
	applicationTokens, err := m.db.ReadApplicationTokensPage(r.Context(), page-1, count)
	if err != nil {
		l.Errorf("db read: %s", err.Error())
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
		return
	}
	for _, at := range applicationTokens {
		if at.CreatedBy == nil {
			creator, err := m.db.ReadFediAccount(r.Context(), at.CreatedByID)
			if err != nil {
				l.Errorf("db read fedi account %d: %s", at.CreatedByID, err.Error())
				m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
				return
			}
			if creator.Instance == nil {
				ownerInstance, err := m.db.ReadFediInstance(r.Context(), creator.InstanceID)
				if err != nil {
					l.Errorf("db read fedi instasnce %d: %s", creator.InstanceID, err.Error())
					m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
					return
				}
				creator.Instance = ownerInstance
			}
			at.CreatedBy = creator
		}
	}
	tmplVars.ApplicationTokens = &applicationTokens

	// count oauth clients
	applicationTokenCount, err := m.db.CountApplicationTokens(r.Context())
	if err != nil {
		l.Errorf("db count: %s", err.Error())
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
		return
	}

	// make pagination
	pageConf := &libtemplate.PaginationConfig{
		Count:         int(applicationTokenCount),
		DisplayCount:  count,
		HRef:          path.AdminSystemApplicationTokens,
		MaxPagination: 5,
		Page:          page,
	}
	if countFound {
		pageConf.HRefCount = count
	}
	tmplVars.Pagination = libtemplate.MakePagination(pageConf)

	err = m.executeTemplate(w, template.AdminSystemApplicationTokensName, tmplVars)
	if err != nil {
		l.Errorf("could not render %s template: %s", template.AdminSystemApplicationTokensName, err.Error())
	}
}
