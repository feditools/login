package webapp

import (
	"github.com/feditools/go-lib/language"
	"github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/path"
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
		HrefAddClient: path.AdminOauthClientAdd,
	}

	// make admin navbar
	navbar := makeAdminNavbar(r)
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

// AdminOauthClientAddGetHandler serves the admin add client page
func (m *Module) AdminOauthClientAddGetHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	m.displayOauthClientAdd(w, r, "", "")
}

// AdminOauthClientAddPostHandler handles the admin add client form
func (m *Module) AdminOauthClientAddPostHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	l := logger.WithField("func", "AdminOauthClientAddGetHandler")

	// get localizer
	localizer := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer)

	// parse form
	err := r.ParseForm()
	if err != nil {
		m.returnErrorPage(w, r, nethttp.StatusBadRequest, err.Error())
	}

	// Init template variables
	tmplVars := &template.AdminOauthClientAdd{
		Common: template.Common{
			PageTitle: "Admin Add Clients",
		},
		Admin: template.Admin{
			Sidebar: makeAdminOauthSidebar(r),
		},

		FormInputDescriptionDisabled: false,
		FormInputDescriptionValue:    "",
		FormButtonSubmitText:         localizer.TextCreate().String(),
	}

	// make admin navbar
	navbar := makeAdminNavbar(r)
	tmplVars.SetNavbar(navbar)

	err = m.initTemplate(w, r, tmplVars)
	if err != nil {
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
		return
	}

	err = m.executeTemplate(w, template.AdminOauthClientAddName, tmplVars)
	if err != nil {
		l.Errorf("could not render %s template: %s", template.AdminOauthClientAddName, err.Error())
	}
}

func (m *Module) displayOauthClientAdd(w nethttp.ResponseWriter, r *nethttp.Request, description, returnURI string) {
	l := logger.WithField("func", "displayOauthClientAdd")

	// get localizer
	localizer := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer)

	// Init template variables
	tmplVars := &template.AdminOauthClientAdd{
		Common: template.Common{
			PageTitle: localizer.TextAddOauth20Client(1).String(),
		},
		Admin: template.Admin{
			Sidebar: makeAdminOauthSidebar(r),
		},

		FormInputDescriptionDisabled: false,
		FormInputDescriptionValue:    description,
		FormInputReturnURIDisabled:   false,
		FormInputReturnURIValue:      returnURI,

		FormInputDescription: &template.FormInput{
			ID:           "inputDescription",
			Type:         "text",
			Name:         "description",
			Placeholder:  localizer.TextDescription(1).String(),
			Label:        localizer.TextDescription(1),
			LabelClass:   "col-sm-2 col-form-label",
			Value:        description,
			WrappedClass: "col-sm-10",
			Disabled:     false,
			Required:     true,
		},

		FormInputReturnURI: &template.FormInput{
			ID:           "inputReturnURI",
			Type:         "text",
			Name:         "return-uri",
			Placeholder:  localizer.TextReturnURI().String(),
			Label:        localizer.TextReturnURI(),
			LabelClass:   "col-sm-2 col-form-label",
			Value:        returnURI,
			WrappedClass: "col-sm-10",
			Disabled:     false,
			Required:     true,
		},

		FormButtonSubmitText: localizer.TextCreate().String(),
	}

	// make admin navbar
	navbar := makeAdminNavbar(r)
	tmplVars.SetNavbar(navbar)

	err := m.initTemplate(w, r, tmplVars)
	if err != nil {
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
		return
	}

	err = m.executeTemplate(w, template.AdminOauthClientAddName, tmplVars)
	if err != nil {
		l.Errorf("could not render %s template: %s", template.AdminOauthClientAddName, err.Error())
	}
}
