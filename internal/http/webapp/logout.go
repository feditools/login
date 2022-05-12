package webapp

import (
	"github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/path"
	"github.com/gorilla/sessions"
	nethttp "net/http"
)

// LogoutGetHandler logs a user out
func (m *Module) LogoutGetHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	// Init Session
	us := r.Context().Value(http.ContextKeySession).(*sessions.Session)

	// Set account to nil
	us.Values[SessionKeyAccountID] = nil
	err := us.Save(r, w)
	if err != nil {
		nethttp.Error(w, err.Error(), nethttp.StatusInternalServerError)
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
		return
	}

	nethttp.Redirect(w, r, path.Login, nethttp.StatusFound)
}
