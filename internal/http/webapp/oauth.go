package webapp

import (
	"errors"
	nethttp "net/http"
	"strconv"

	"github.com/feditools/login/internal/models"

	"github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/path"
	oerrors "github.com/go-oauth2/oauth2/v4/errors"
	"github.com/gorilla/sessions"
)

// OauthAuthorizeGetHandler handles oauth authorization.
func (m *Module) OauthAuthorizeGetHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	l := logger.WithField("func", "OauthAuthorizeGetHandler")

	_, shouldReturn := m.authRequireLoggedIn(w, r)
	if shouldReturn {
		return
	}

	// get session
	user := r.Context().Value(http.ContextKeyAccount).(*models.FediAccount)

	err := m.oauth.HandleAuthorizeRequest(user.ID, w, r)
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

// OauthTokenHandler handles oauth tokens.
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

	uid, ok := us.Values[http.SessionKeyAccountID].(int64)
	if !ok {
		if r.Form == nil {
			err := r.ParseForm()
			if err != nil {
				l.Errorf("parsing form: %s", err.Error())

				return "", err
			}
		}

		// Save current page
		us.Values[http.SessionKeyReturnURI] = r.Form
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
