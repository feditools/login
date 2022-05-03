package webapp

import (
	"github.com/feditools/go-lib/language"
	"github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/path"
	"github.com/feditools/login/internal/template"
	nethttp "net/http"
	"net/url"
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
	m.displayOauthClientAdd(w, r, "", "", nil, nil)
}

// AdminOauthClientAddPostHandler handles the admin add client form
func (m *Module) AdminOauthClientAddPostHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	l := logger.WithField("func", "AdminOauthClientAddPostHandler")

	// get localizer
	localizer := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer)

	// parse form
	err := r.ParseForm()
	if err != nil {
		m.returnErrorPage(w, r, nethttp.StatusBadRequest, err.Error())
	}

	// get vars and validate
	valid := true
	descriptionValidation := &template.FormValidation{
		Valid:    true,
		Response: localizer.TextLooksGood().String(),
	}
	description := r.Form.Get(FormDescription)
	if description == "" {
		descriptionValidation = &template.FormValidation{
			Valid:    false,
			Response: localizer.TextLooksGood().String(),
		}
		valid = false
	}
	returnURIValidation := &template.FormValidation{
		Valid:    true,
		Response: localizer.TextLooksGood().String(),
	}
	returnURI := r.Form.Get(FormReturnURI)
	if returnURI == "" {
		returnURIValidation = &template.FormValidation{
			Valid:    false,
			Response: localizer.TextLooksGood().String(),
		}
		valid = false
	}
	_, err = url.Parse(returnURI)
	if err != nil {
		returnURIValidation = &template.FormValidation{
			Valid:    false,
			Response: localizer.TextLooksGood().String(),
		}
		valid = false
	}

	// return form if invalid
	if !valid {
		m.displayOauthClientAdd(w, r, description, returnURI, descriptionValidation, returnURIValidation)
		return
	}

	// Init template variables
	tmplVars := &template.AdminOauthClientAdded{
		Common: template.Common{
			PageTitle: localizer.TextAddOauth20Client(1).String(),
		},
		Admin: template.Admin{
			Sidebar: makeAdminOauthSidebar(r),
		},

		FormInputDescription: &template.FormInput{
			ID:           "inputDescription",
			Type:         "text",
			Name:         FormDescription,
			Placeholder:  localizer.TextDescription(1).String(),
			Label:        localizer.TextDescription(1),
			LabelClass:   "col-sm-2 col-form-label",
			Value:        "",
			WrappedClass: "col-sm-10",
			Disabled:     true,
			Required:     true,
		},
		FormInputClientID: &template.FormInput{
			ID:           "inputClientID",
			Type:         "text",
			Name:         "client-id",
			Placeholder:  localizer.TextClientID(1).String(),
			Label:        localizer.TextClientID(1),
			LabelClass:   "col-sm-2 col-form-label",
			Value:        "",
			WrappedClass: "col-sm-10",
			Disabled:     true,
			Required:     true,
		},
		FormInputClientSecret: &template.FormInput{
			ID:           "inputClientSecret",
			Type:         "text",
			Name:         "client-secret",
			Placeholder:  localizer.TextClientSecret(1).String(),
			Label:        localizer.TextClientSecret(1),
			LabelClass:   "col-sm-2 col-form-label",
			Value:        "",
			WrappedClass: "col-sm-10",
			Disabled:     true,
			Required:     true,
		},
		FormInputReturnURI: &template.FormInput{
			ID:           "inputReturnURI",
			Type:         "text",
			Name:         "return-uri",
			Placeholder:  localizer.TextReturnURI().String(),
			Label:        localizer.TextReturnURI(),
			LabelClass:   "col-sm-2 col-form-label",
			Value:        "",
			WrappedClass: "col-sm-10",
			Disabled:     true,
			Required:     true,
		},
	}

	// make admin navbar
	navbar := makeAdminNavbar(r)
	tmplVars.SetNavbar(navbar)

	err = m.initTemplate(w, r, tmplVars)
	if err != nil {
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
		return
	}

	err = m.executeTemplate(w, template.AdminOauthClientAddedName, tmplVars)
	if err != nil {
		l.Errorf("could not render %s template: %s", template.AdminOauthClientAddName, err.Error())
	}
}

func (m *Module) displayOauthClientAdd(w nethttp.ResponseWriter, r *nethttp.Request, description, returnURI string, descriptionVal, returnURIVal *template.FormValidation) {
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

		FormInputDescription: &template.FormInput{
			ID:           "inputDescription",
			Type:         "text",
			Name:         FormDescription,
			Placeholder:  localizer.TextDescription(1).String(),
			Label:        localizer.TextDescription(1),
			LabelClass:   "col-sm-2 col-form-label",
			Value:        description,
			WrappedClass: "col-sm-10",
			Disabled:     false,
			Required:     true,
			Validation:   descriptionVal,
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
			Validation:   returnURIVal,
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
