package webapp

import (
	"github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/language"
	"github.com/feditools/login/internal/path"
	"github.com/feditools/login/internal/template"
	nethttp "net/http"
)

// LoginGetHandler serves the login page
func (m *Module) LoginGetHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	m.displayLoginPage(w, r, "", "")
}

// LoginPostHandler attempts a login
func (m *Module) LoginPostHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	l := logger.WithField("func", "LoginPostHandler")

	// get localizer
	//localizer := r.Context().Value(localizerContextKey).(*language.Localizer)

	// parse form data
	err := r.ParseForm()
	if err != nil {
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
		return
	}

	// check if account exists
	formAccount := r.Form.Get("account")
	loginURL, err := m.fedi.GetLoginURL(r.Context(), formAccount)
	if err != nil {
		l.Errorf("get login url: %s", err.Error())
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
		return
	}

	l.Debugf("login url: %s", loginURL.String())
	nethttp.Redirect(w, r, loginURL.String(), nethttp.StatusFound)
	return
}

func (m *Module) displayLoginPage(w nethttp.ResponseWriter, r *nethttp.Request, account, formError string) {
	l := logger.WithField("func", "displayLoginPage")

	// get localizer
	localizer := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer)

	// Init template variables
	tmplVars := &template.Login{}
	err := m.initTemplate(w, r, tmplVars)
	if err != nil {
		nethttp.Error(w, err.Error(), nethttp.StatusInternalServerError)
		return
	}

	// add error css file
	signature, err := m.getSignatureCached(DirWeb + path.FileLoginCSS)
	if err != nil {
		l.Errorf("getting signature for %s: %s", path.FileLoginCSS, err.Error())
	}
	tmplVars.AddHeadLink(template.HeadLink{
		HRef:        path.FileLoginCSS,
		Rel:         "stylesheet",
		CrossOrigin: "anonymous",
		Integrity:   signature,
	})

	tmplVars.PageTitle = localizer.TextLogin().String()

	// set bot image
	tmplVars.Image = m.logoURI

	// set form values
	tmplVars.FormError = formError
	tmplVars.FormAccount = account

	err = m.executeTemplate(w, "login", tmplVars)
	if err != nil {
		l.Errorf("could not render login template: %s", err.Error())
	}
}
