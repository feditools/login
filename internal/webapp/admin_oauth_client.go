package webapp

import (
	"github.com/feditools/go-lib/language"
	libtemplate "github.com/feditools/go-lib/template"
	"github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/models"
	"github.com/feditools/login/internal/path"
	"github.com/feditools/login/internal/template"
	"github.com/google/uuid"
	"mvdan.cc/xurls/v2"
	nethttp "net/http"
)

// AdminOauthClientsGetHandler serves the admin client page
func (m *Module) AdminOauthClientsGetHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	l := logger.WithField("func", "AdminOauthClientsGetHandler")

	// get localizer
	localizer := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer)

	// Init template variables
	tmplVars := &template.AdminOauthClients{
		Common: template.Common{
			PageTitle: localizer.TextOauth20Client(2).String(),
		},
		Admin: template.Admin{
			Sidebar: makeAdminOauthSidebar(r),
		},
		HrefAddClient: path.AdminOauthClientAdd,
	}

	err := m.initTemplateAdmin(w, r, tmplVars)
	if err != nil {
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
		return
	}

	err = m.executeTemplate(w, template.AdminOauthClientsName, tmplVars)
	if err != nil {
		l.Errorf("could not render %s template: %s", template.AdminOauthClientsName, err.Error())
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
	descriptionValidation := &libtemplate.FormValidation{
		Valid:    true,
		Response: localizer.TextLooksGood().String(),
	}
	description := r.Form.Get(FormDescription)
	if description == "" {
		descriptionValidation = &libtemplate.FormValidation{
			Valid:    false,
			Response: localizer.TextRequired().String(),
		}
		valid = false
	}
	redirectURIValidation := &libtemplate.FormValidation{
		Valid:    true,
		Response: localizer.TextLooksGood().String(),
	}
	redirectURI := r.Form.Get(FormRedirectURI)
	if redirectURI == "" {
		redirectURIValidation = &libtemplate.FormValidation{
			Valid:    false,
			Response: localizer.TextRequired().String(),
		}
		valid = false
	}
	if redirectURIValidation.Valid {
		rxStrict, err := xurls.StrictMatchingScheme("(http|https)")
		if err != nil {
			l.Warnf("couldn't compile regex")
		}
		matches := rxStrict.FindAllString(redirectURI, -1)
		if len(matches) != 1 {
			// url not found or too many uris
			redirectURIValidation = &libtemplate.FormValidation{
				Valid:    false,
				Response: localizer.TextInvalidURI(1).String(),
			}
			valid = false
		} else if matches[0] != redirectURI {
			// check for extraneous text
			redirectURIValidation = &libtemplate.FormValidation{
				Valid:    false,
				Response: localizer.TextInvalidURI(1).String(),
			}
			valid = false
		}
	}

	// return form if invalid
	if !valid {
		m.displayOauthClientAdd(w, r, description, redirectURI, descriptionValidation, redirectURIValidation)
		return
	}

	// get account
	account := r.Context().Value(http.ContextKeyAccount).(*models.FediAccount)

	// add to the database
	newClientSecret := uuid.New().String()
	newClient := &models.OauthClient{
		Description: description,
		RedirectURI: redirectURI,
		OwnerID:     account.ID,
	}
	err = newClient.SetSecret(newClientSecret)
	if err != nil {
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
		return
	}
	err = m.db.CreateOauthClient(r.Context(), newClient)
	if err != nil {
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
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

		FormInputDescription: &libtemplate.FormInput{
			ID:           "inputDescription",
			Type:         "text",
			Name:         FormDescription,
			Placeholder:  localizer.TextDescription(1).String(),
			Label:        localizer.TextDescription(1),
			LabelClass:   "col-sm-2 col-form-label",
			Value:        newClient.Description,
			WrappedClass: "col-sm-10",
			Disabled:     true,
			Required:     true,
		},
		FormInputClientID: &libtemplate.FormInput{
			ID:           "inputClientID",
			Type:         "text",
			Name:         "client-id",
			Placeholder:  localizer.TextClientID(1).String(),
			Label:        localizer.TextClientID(1),
			LabelClass:   "col-sm-2 col-form-label",
			Value:        m.tokenizer.GetToken(newClient),
			WrappedClass: "col-sm-10",
			Disabled:     true,
			Required:     true,
		},
		FormInputClientSecret: &libtemplate.FormInput{
			ID:           "inputClientSecret",
			Type:         "text",
			Name:         "client-secret",
			Placeholder:  localizer.TextClientSecret(1).String(),
			Label:        localizer.TextClientSecret(1),
			LabelClass:   "col-sm-2 col-form-label",
			Value:        newClientSecret,
			WrappedClass: "col-sm-10",
			Disabled:     true,
			Required:     true,
		},
		FormInputRedirectURI: &libtemplate.FormInput{
			ID:           "inputRedirectURI",
			Type:         "text",
			Name:         FormRedirectURI,
			Placeholder:  localizer.TextRedirectURI(1).String(),
			Label:        localizer.TextRedirectURI(1),
			LabelClass:   "col-sm-2 col-form-label",
			Value:        newClient.RedirectURI,
			WrappedClass: "col-sm-10",
			Disabled:     true,
			Required:     true,
		},
	}

	err = m.initTemplateAdmin(w, r, tmplVars)
	if err != nil {
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
		return
	}

	err = m.executeTemplate(w, template.AdminOauthClientAddedName, tmplVars)
	if err != nil {
		l.Errorf("could not render %s template: %s", template.AdminOauthClientAddName, err.Error())
	}
}

func (m *Module) displayOauthClientAdd(w nethttp.ResponseWriter, r *nethttp.Request, description, redirectURI string, descriptionVal, redirectURIVal *libtemplate.FormValidation) {
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

		FormInputDescription: &libtemplate.FormInput{
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
		FormInputRedirectURI: &libtemplate.FormInput{
			ID:           "inputRedirectURI",
			Type:         "text",
			Name:         FormRedirectURI,
			Placeholder:  localizer.TextRedirectURI(1).String(),
			Label:        localizer.TextRedirectURI(1),
			LabelClass:   "col-sm-2 col-form-label",
			Value:        redirectURI,
			WrappedClass: "col-sm-10",
			Disabled:     false,
			Required:     true,
			Validation:   redirectURIVal,
		},

		FormButtonSubmitText: localizer.TextCreate().String(),
	}

	err := m.initTemplateAdmin(w, r, tmplVars)
	if err != nil {
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
		return
	}

	err = m.executeTemplate(w, template.AdminOauthClientAddName, tmplVars)
	if err != nil {
		l.Errorf("could not render %s template: %s", template.AdminOauthClientAddName, err.Error())
	}
}
