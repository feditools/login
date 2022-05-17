package webapp

import (
	nethttp "net/http"

	libhttp "github.com/feditools/go-lib/http"
	"github.com/feditools/go-lib/language"
	libtemplate "github.com/feditools/go-lib/template"
	"github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/http/template"
	"github.com/feditools/login/internal/models"
	"github.com/feditools/login/internal/path"
	"github.com/google/uuid"
)

// AdminSystemApplicationTokensGetHandler serves the admin client page.
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

	page, count, countFound := libhttp.GetPaginationFromURL(r.URL, defaultCount)

	// get application tokens
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

// AdminSystemApplicationTokenAddGetHandler serves the admin add application token page.
func (m *Module) AdminSystemApplicationTokenAddGetHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	m.displayApplicationTokenAdd(w, r, "", nil)
}

// AdminSystemApplicationTokenAddPostHandler handles the admin add client form.
func (m *Module) AdminSystemApplicationTokenAddPostHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	l := logger.WithField("func", "AdminSystemApplicationTokenAddPostHandler")

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

	// return form if invalid
	if !valid {
		m.displayApplicationTokenAdd(w, r, description, descriptionValidation)
		return
	}

	// get account
	account := r.Context().Value(http.ContextKeyAccount).(*models.FediAccount)

	// add to the database
	newApplicationToken := &models.ApplicationToken{
		Description: description,
		Token:       uuid.New().String(),
		CreatedByID: account.ID,
	}
	err = m.db.CreateApplicationToken(r.Context(), newApplicationToken)
	if err != nil {
		l.Errorf("db craete: %s", err.Error())
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
		return
	}

	m.displayApplicationTokenAdded(w, r, description, newApplicationToken.Token)
}

func (m *Module) displayApplicationTokenAdd(w nethttp.ResponseWriter, r *nethttp.Request, description string, descriptionVal *libtemplate.FormValidation) {
	l := logger.WithField("func", "displayApplicationTokenAdd")

	// get localizer
	localizer := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer)

	// Init template variables
	tmplVars := &template.AdminSystemApplicationTokenAdd{
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

		FormButtonSubmitText: localizer.TextCreate().String(),
	}

	err := m.initTemplateAdmin(w, r, tmplVars)
	if err != nil {
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
		return
	}

	err = m.executeTemplate(w, template.AdminSystemApplicationTokenAddName, tmplVars)
	if err != nil {
		l.Errorf("could not render %s template: %s", template.AdminSystemApplicationTokenAddName, err.Error())
	}
}

func (m *Module) displayApplicationTokenAdded(w nethttp.ResponseWriter, r *nethttp.Request, description, token string) {
	l := logger.WithField("func", "displayApplicationTokenAdd")

	// get localizer
	localizer := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer)

	// Init template variables
	tmplVars := &template.AdminSystemApplicationTokenAdded{
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
			Disabled:     true,
			Required:     true,
		},
		FormInputToken: &libtemplate.FormInput{
			ID:           "inputToken",
			Type:         "text",
			Name:         FormToken,
			Placeholder:  localizer.TextToken(1).String(),
			Label:        localizer.TextToken(1),
			LabelClass:   "col-sm-2 col-form-label",
			Value:        token,
			WrappedClass: "col-sm-10",
			Disabled:     true,
			Required:     true,
		},
	}

	err := m.initTemplateAdmin(w, r, tmplVars)
	if err != nil {
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
		return
	}

	err = m.executeTemplate(w, template.AdminSystemApplicationTokenAddedName, tmplVars)
	if err != nil {
		l.Errorf("could not render %s template: %s", template.AdminSystemApplicationTokenAddedName, err.Error())
	}
}
