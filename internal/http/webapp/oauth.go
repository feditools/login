package webapp

import (
	"errors"
	"github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/path"
	oerrors "github.com/go-oauth2/oauth2/v4/errors"
	"github.com/gorilla/sessions"
	nethttp "net/http"
	"strconv"
)

// OauthAuthorizeGetHandler handles oauth authorization
func (m *Module) OauthAuthorizeGetHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	l := logger.WithField("func", "OauthAuthorizeGetHandler")

	_, shouldReturn := m.authRequireLoggedIn(w, r)
	if shouldReturn {
		return
	}

	err := m.oauth.HandleAuthorizeRequest(w, r)
	if err != nil {
		switch {
		case errors.Is(err, oerrors.ErrCodeChallengeRquired):
			l.Debugf("ErrCodeChallengeRquired: %s", err.Error())
		case errors.Is(err, oerrors.ErrUnsupportedCodeChallengeMethod):
			l.Debugf("ErrUnsupportedCodeChallengeMethod: %s", err.Error())
		case errors.Is(err, oerrors.ErrInvalidRequest):
			l.Debugf("ErrInvalidRequest: %s", err.Error())
		default:
			l.Debugf("unknown: %s", err.Error())
		}
		m.returnErrorPage(w, r, nethttp.StatusBadRequest, err.Error())
	}
}

// OauthTokenHandler handles oauth tokens
func (m *Module) OauthTokenHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	l := logger.WithField("func", "OauthTokenHandler")
	dumpRequest(l, "token", r)

	err := m.oauth.HandleTokenRequest(w, r)
	if err != nil {
		l.Errorf("error: %s", err.Error())
	}
}

func oauthUserAuthorizeHandler(w nethttp.ResponseWriter, r *nethttp.Request) (string, error) {
	l := logger.WithField("func", "oauthUserAuthorizeHandler")

	// get session
	us := r.Context().Value(http.ContextKeySession).(*sessions.Session)

	uid, ok := us.Values[SessionKeyAccountID].(int64)
	if !ok {
		if r.Form == nil {
			r.ParseForm()
		}

		// Save current page
		us.Values[SessionKeyReturnURI] = r.Form
		err := us.Save(r, w)
		if err != nil {
			l.Errorf("saving session: %s", err.Error())
			return "", err
		}

		w.Header().Set("Location", path.Login)
		w.WriteHeader(nethttp.StatusFound)
		return "", nil
	}

	return strconv.FormatInt(uid, 10), nil
}
