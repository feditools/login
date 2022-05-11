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

// AdminFediInstancesGetHandler serves the admin client page
func (m *Module) AdminFediInstancesGetHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	l := logger.WithField("func", "AdminFediInstancesGetHandler")

	// get localizer
	localizer := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer)

	// Init template variables
	tmplVars := &template.AdminFediInstances{
		Common: template.Common{
			PageTitle: localizer.TextOauth20Client(2).String(),
		},
		Admin: template.Admin{
			Sidebar: makeAdminFediverseSidebar(r),
		},
		HRefViewFediInstance: path.AdminFediverseInstances,
	}

	err := m.initTemplateAdmin(w, r, tmplVars)
	if err != nil {
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
		return
	}

	page, count, countFound := lib.GetPaginationFromURL(r.URL, defaultCount)

	// get oauth clients
	instances, err := m.db.ReadFediInstancesPage(r.Context(), page-1, count)
	if err != nil {
		l.Errorf("db read: %s", err.Error())
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
		return
	}
	instanceAccountCounts := make([]int64, len(instances))
	for i, c := range instances {
		accountCount, err := m.db.CountFediAccountsForInstance(r.Context(), c.ID)
		if err != nil {
			l.Errorf("db count: %s", err.Error())
			m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
			return
		}
		instanceAccountCounts[i] = accountCount
	}
	tmplVars.FediInstances = &instances
	tmplVars.FediInstanceAccountCounts = instanceAccountCounts

	// count oauth clients
	instanceCount, err := m.db.CountFediInstances(r.Context())
	if err != nil {
		l.Errorf("db count: %s", err.Error())
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
		return
	}

	// make pagination
	pageConf := &libtemplate.PaginationConfig{
		Count:         int(instanceCount),
		DisplayCount:  count,
		HRef:          path.AdminFediverseInstances,
		MaxPagination: 5,
		Page:          page,
	}
	if countFound {
		pageConf.HRefCount = count
	}
	tmplVars.Pagination = libtemplate.MakePagination(pageConf)

	err = m.executeTemplate(w, template.AdminFediInstancesName, tmplVars)
	if err != nil {
		l.Errorf("could not render %s template: %s", template.AdminFediInstancesName, err.Error())
	}
}
