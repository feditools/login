package webapp

import (
	"context"
	"github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/language"
	"github.com/go-http-utils/etag"
	nethttp "net/http"
)

// Middleware runs on every http request
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
		if accountID, ok := val.(int64); ok {
			// read federated accounts
			account, err := m.db.ReadFediAccount(ctx, accountID)
			if err != nil {
				l.Errorf("db read: %s", err.Error())
				m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
				return
			}

			if account != nil {
				// read federated instance
				instance, err := m.db.ReadFediInstance(ctx, account.InstanceID)
				if err != nil {
					l.Errorf("db read: %s", err.Error())
					m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
					return
				}
				account.Instance = instance

				ctx = context.WithValue(ctx, http.ContextKeyAccount, account)
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

// MiddlewareRequireAdmin will redirect a user to login page if user not in context and will return unauthorized for
// a non admin user.
func (m *Module) MiddlewareRequireAdmin(next nethttp.Handler) nethttp.Handler {
	return nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		account, shouldReturn := m.authRequireLoggedIn(w, r)
		if shouldReturn {
			return
		}

		if !account.Admin {
			localizer := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer)
			m.returnErrorPage(w, r, nethttp.StatusUnauthorized, localizer.TextUnauthorized().String())
			return
		}

		next.ServeHTTP(w, r)
	})
}
