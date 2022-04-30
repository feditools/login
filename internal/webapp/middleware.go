package webapp

import (
	"context"
	"github.com/feditools/login/internal/http"
	"github.com/go-http-utils/etag"
	"golang.org/x/text/language"
	nethttp "net/http"
)

// Middleware runs on every nethttp request
func (m *Module) Middleware(next nethttp.Handler) nethttp.Handler {
	return etag.Handler(nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		l := logger.WithField("func", "Middleware")

		// Init Session
		us, err := m.store.Get(r, "login")
		if err != nil {
			l.Errorf("get session: %s", err.Error())
			m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
			return
		}
		ctx := context.WithValue(r.Context(), http.ContextKeySession, us)

		// Retrieve our account and type-assert it
		val := us.Values[SessionKeyAccountID]
		if userID, ok := val.(int64); ok {
			// read federated accounts
			user, err := m.db.ReadFediAccount(ctx, userID)
			if err != nil {
				l.Errorf("db read: %s", err.Error())
				m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
				return
			}

			if user != nil {
				// read federated instance
				instance, err := m.db.ReadFediInstance(ctx, user.InstanceID)
				if err != nil {
					l.Errorf("db read: %s", err.Error())
					m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
					return
				}
				user.Instance = instance

				ctx = context.WithValue(ctx, http.ContextKeyAccount, user)
			}
		}

		// create localizer
		lang := r.FormValue("lang")
		accept := r.Header.Get("Accept-Language")
		localizer, err := m.language.NewLocalizer(lang, accept)
		if err != nil {
			l.Errorf("could get localizer: %s", err.Error())
			m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
			return
		}
		ctx = context.WithValue(ctx, http.ContextKeyLocalizer, localizer)

		// set request language
		ctx = context.WithValue(ctx, http.ContextKeyLanguage, getPageLang(lang, accept, m.language.Language().String()))

		// Do Request
		next.ServeHTTP(w, r.WithContext(ctx))
	}), false)
}

func getPageLang(query, header, defaultLang string) string {
	l := logger.WithField("func", "getPageLang")

	if query != "" {
		t, _, err := language.ParseAcceptLanguage(query)
		if err == nil {
			l.Debugf("query languages: %v", t)
			if len(t) > 0 {
				l.Debugf("returning language: %s", t[0].String())
				return t[0].String()
			}
		} else {
			l.Debugf("query '%s' did not contain a valid lanaugae: %s", query, err.Error())
		}
	}

	if header != "" {
		t, _, err := language.ParseAcceptLanguage(header)
		if err == nil {
			l.Debugf("header languages: %v", t)
			if len(t) > 0 {
				l.Debugf("returning language: %s", t[0].String())
				return t[0].String()
			}
		} else {
			l.Debugf("query '%s' did not contain a valid lanaugae: %s", query, err.Error())
		}
	}

	l.Debugf("returning default language: %s", defaultLang)
	return defaultLang
}
