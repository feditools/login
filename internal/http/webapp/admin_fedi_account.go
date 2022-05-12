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

// AdminFediAccountsGetHandler serves the admin client page
func (m *Module) AdminFediAccountsGetHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	l := logger.WithField("func", "AdminFediAccountsGetHandler")

	// get localizer
	localizer := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer)

	// Init template variables
	tmplVars := &template.AdminFediAccounts{
		Common: template.Common{
			PageTitle: localizer.TextOauth20Client(2).String(),
		},
		Admin: template.Admin{
			Sidebar: makeAdminFediverseSidebar(r),
		},
		HRefViewFediAccount:  path.AdminFediverseAccounts,
		HRefViewFediInstance: path.AdminFediverseInstances,
	}

	err := m.initTemplateAdmin(w, r, tmplVars)
	if err != nil {
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
		return
	}

	page, count, countFound := lib.GetPaginationFromURL(r.URL, defaultCount)

	// get oauth clients
	accounts, err := m.db.ReadFediAccountsPage(r.Context(), page-1, count)
	if err != nil {
		l.Errorf("db read: %s", err.Error())
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
		return
	}
	for _, a := range accounts {
		if a.Instance == nil {
			instance, err := m.db.ReadFediInstance(r.Context(), a.InstanceID)
			if err != nil {
				l.Errorf("db read fedi instance %d: %s", a.InstanceID, err.Error())
				m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
				return
			}
			a.Instance = instance
		}
	}
	tmplVars.FediAccounts = &accounts

	// count oauth clients
	accountCount, err := m.db.CountFediAccounts(r.Context())
	if err != nil {
		l.Errorf("db count: %s", err.Error())
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
		return
	}

	// make pagination
	pageConf := &libtemplate.PaginationConfig{
		Count:         int(accountCount),
		DisplayCount:  count,
		HRef:          path.AdminFediverseAccounts,
		MaxPagination: 5,
		Page:          page,
	}
	if countFound {
		pageConf.HRefCount = count
	}
	tmplVars.Pagination = libtemplate.MakePagination(pageConf)

	err = m.executeTemplate(w, template.AdminFediAccountsName, tmplVars)
	if err != nil {
		l.Errorf("could not render %s template: %s", template.AdminFediAccountsName, err.Error())
	}
}
