package webapp

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/language"
	"github.com/feditools/login/internal/models"
	"github.com/feditools/login/internal/template"
	"github.com/gorilla/sessions"
	nethttp "net/http"
)

func (m *Module) initTemplate(w nethttp.ResponseWriter, r *nethttp.Request, tmpl template.InitTemplate) error {
	l := logger.WithField("func", "initTemplate")

	// set text handler
	localizer := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer)
	tmpl.SetLocalizer(localizer)

	// set language
	lang := r.Context().Value(http.ContextKeyLanguage).(string)
	tmpl.SetLanguage(lang)

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
		tmpl.SetAccountID(account.ID)
	}

	// try to read session data
	if r.Context().Value(http.ContextKeySession) == nil {
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
	}

	return nil
}

func (m *Module) executeTemplate(w nethttp.ResponseWriter, name string, tmplVars interface{}) error {
	b := new(bytes.Buffer)
	err := m.templates.ExecuteTemplate(b, name, tmplVars)
	if err != nil {
		return err
	}

	h := sha256.New()
	h.Write(b.Bytes())
	w.Header().Set("Digest", fmt.Sprintf("sha-256=%s", base64.StdEncoding.EncodeToString(h.Sum(nil))))

	if m.minify == nil {
		_, err := w.Write(b.Bytes())
		return err
	}
	return m.minify.Minify("text/html", w, b)
}