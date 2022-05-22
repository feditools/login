package webapp

import (
	"fmt"
	nethttp "net/http"

	"github.com/feditools/login/internal/fedi"
	"github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/models"
	"github.com/feditools/login/internal/path"
	"github.com/feditools/login/internal/token"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

// CallbackOauthGetHandler handles an oauth callback.
func (m *Module) CallbackOauthGetHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	l := logger.WithField("func", "CallbackMastodonGetHandler")

	// lookup instance
	vars := mux.Vars(r)
	kind, id, err := m.tokenizer.DecodeToken(vars[path.VarInstanceID])
	if err != nil {
		l.Debugf("decode token: %s", err.Error())
		m.returnErrorPage(w, r, nethttp.StatusBadRequest, "bad token")

		return
	}
	if kind != token.KindFediInstance {
		l.Debug("token is wrong kind")
		m.returnErrorPage(w, r, nethttp.StatusBadRequest, "bad token")

		return
	}
	instance, err := m.db.ReadFediInstance(r.Context(), id)
	if err != nil {
		l.Errorf("db read instance: %s", err.Error())
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())

		return
	}
	if instance == nil {
		m.returnErrorPage(w, r, nethttp.StatusNotFound, vars["token"])

		return
	}

	switch fedi.Software(instance.Software) {
	case fedi.SoftwareMastodon:
		// get code
		var code []string
		var ok bool
		if code, ok = r.URL.Query()["code"]; !ok || len(code[0]) < 1 {
			l.Debugf("missing code")
			m.returnErrorPage(w, r, nethttp.StatusBadRequest, "missing code")

			return
		}

		// retrieve access token
		var accessToken string
		accessToken, err = m.fedi.Helper(fedi.SoftwareMastodon).GetAccessToken(r.Context(), instance, code[0])
		if err != nil {
			l.Errorf("get access token error: %s", err.Error())
			m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())

			return
		}
		l.Debugf("access token: %s", accessToken)

		// retrieve current account
		var account *models.FediAccount
		account, err = m.fedi.Helper(fedi.SoftwareMastodon).GetCurrentAccount(r.Context(), instance, accessToken)
		if err != nil {
			l.Errorf("get access token error: %s", err.Error())
			m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())

			return
		}

		// increment login
		err = m.db.IncFediAccountLoginCount(r.Context(), account)
		if err != nil {
			l.Errorf("db inc login: %s", err.Error())
			m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())

			return
		}

		// init session
		us := r.Context().Value(http.ContextKeySession).(*sessions.Session)
		us.Values[http.SessionKeyAccountID] = account.ID
		err = us.Save(r, w)
		if err != nil {
			m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())

			return
		}

		l.Debugf("account: %#v", account)

		// redirect to last page
		val := us.Values[http.SessionKeyLoginRedirect]
		var loginRedirect string
		if loginRedirect, ok = val.(string); !ok {
			// redirect home page if no login-redirect
			nethttp.Redirect(w, r, path.Me, nethttp.StatusFound)

			return
		}

		// Set login redirect to nil
		us.Values[http.SessionKeyLoginRedirect] = nil
		err := us.Save(r, w)
		if err != nil {
			m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())

			return
		}
		nethttp.Redirect(w, r, loginRedirect, nethttp.StatusFound)

		return
	default:
		m.returnErrorPage(w, r, nethttp.StatusNotImplemented, fmt.Sprintf("no helper for '%s'", instance.Software))

		return
	}
}
