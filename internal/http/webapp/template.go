package webapp

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	nethttp "net/http"

	"github.com/feditools/go-lib/language"
	"github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/http/template"
	"github.com/feditools/login/internal/models"
	"github.com/feditools/login/internal/path"
)

func (m *Module) initTemplate(_ nethttp.ResponseWriter, r *nethttp.Request, tmpl template.InitTemplate) error {
	// l := logger.WithField("func", "initTemplate")

	// set text handler
	localizer := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer)
	tmpl.SetLocalizer(localizer)

	// set language
	lang := r.Context().Value(http.ContextKeyLanguage).(string)
	tmpl.SetLanguage(lang)

	// set logo image src
	tmpl.SetLogoSrc(m.logoSrcDark, m.logoSrcLight)

	// add css
	for _, link := range m.headLinks {
		tmpl.AddHeadLink(link)
	}

	// add scripts
	for _, script := range m.footerScripts {
		tmpl.AddFooterScript(script)
	}

	if r.Context().Value(http.ContextKeyAccount) != nil {
		account := r.Context().Value(http.ContextKeyAccount).(*models.FediAccount)
		tmpl.SetAccount(account)
	}

	// try to read session data
	/*if r.Context().Value(http.ContextKeySession) == nil {
		return nil
	}

	us := r.Context().Value(http.ContextKeySession).(*sessions.Session)
	saveSession := false

	if saveSession {
		err := us.Save(r, w)
		if err != nil {
			l.Warningf("initTemplate could not save session: %s", err.Error())
			return err
		}
	}*/

	return nil
}

func (m *Module) initTemplateAdmin(w nethttp.ResponseWriter, r *nethttp.Request, tmpl template.InitTemplate) error {
	err := m.initTemplate(w, r, tmpl)
	if err != nil {
		return err
	}

	// make admin navbar
	navbar := makeAdminNavbar(r)
	tmpl.SetNavbar(navbar)

	return nil
}

func (m *Module) executeTemplate(w nethttp.ResponseWriter, name string, tmplVars interface{}) error {
	l := logger.WithField("func", "executeTemplate")

	b := new(bytes.Buffer)
	err := m.templates.ExecuteTemplate(b, name, tmplVars)
	if err != nil {
		return err
	}

	h := sha256.New()
	_, err = h.Write(b.Bytes())
	if err != nil {
		l.Errorf("writing response: %s", err.Error())

		return err
	}
	w.Header().Set("Digest", fmt.Sprintf("sha-256=%s", base64.StdEncoding.EncodeToString(h.Sum(nil))))

	if m.minify == nil {
		_, err := w.Write(b.Bytes())

		return err
	}

	return m.minify.Minify("text/html", w, b)
}

func makeAdminNavbar(r *nethttp.Request) template.Navbar {
	// get localizer
	l := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer)

	// create navbar
	newNavbar := template.Navbar{
		{
			Text:     l.TextHomeWeb().String(),
			MatchStr: path.ReAdmin,
			FAIcon:   "home",
			URL:      path.Admin,
		},
		{
			Text:     l.TextFediverse().String(),
			MatchStr: path.ReAdminFediversePre,
			FAIcon:   "home",
			URL:      path.AdminFediverse,
		},
		{
			Text:     l.TextOauth().String(),
			MatchStr: path.ReAdminOauthPre,
			FAIcon:   "desktop",
			URL:      path.AdminOauth,
		},
		{
			Text:     l.TextSystem(1).String(),
			MatchStr: path.ReAdminSystemPre,
			FAIcon:   "desktop",
			URL:      path.AdminSystem,
		},
	}

	newNavbar.ActivateFromPath(r.URL.Path)

	return newNavbar
}
