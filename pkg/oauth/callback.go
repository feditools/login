package oauth

import (
	"fmt"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	nethttp "net/http"
)

// HandleCallback handles the callback from the oauth server.
func (c *Client) HandleCallback(w nethttp.ResponseWriter, r *nethttp.Request, us *sessions.Session, sessionID string) (*oauth2.Token, error) {
	expectedCode, ok := us.Values[SessionKeyCode].(string)
	if !ok {
		return nil, NewError(nethttp.StatusBadRequest, "missing code")
	}
	expectedNonce, ok := us.Values[SessionKeyNonce].(string)
	if !ok {
		return nil, NewError(nethttp.StatusBadRequest, "missing nonce")
	}
	expectedState, ok := us.Values[SessionKeyState].(string)
	if !ok {
		return nil, NewError(nethttp.StatusBadRequest, "missing state")
	}

	// delete so code and state can't be reused
	us.Values[SessionKeyCode] = nil
	us.Values[SessionKeyState] = nil
	err := us.Save(r, w)
	if err != nil {
		return nil, NewError(nethttp.StatusInternalServerError, fmt.Sprintf("session: %s", err.Error()))
	}

	// parse form
	if err := r.ParseForm(); err != nil {
		return nil, NewError(nethttp.StatusInternalServerError, fmt.Sprintf("form: %s", err.Error()))
	}

	// compare state
	if state := r.Form.Get("state"); state != expectedState {
		return nil, NewError(nethttp.StatusBadRequest, "state invalid")
	}

	// get code
	code := r.Form.Get("code")
	if code == "" {
		return nil, NewError(nethttp.StatusBadRequest, "code not found")
	}

	// request token
	token, err := c.config.Exchange(
		r.Context(),
		code,
		oauth2.SetAuthURLParam("session_id", sessionID),
		oauth2.SetAuthURLParam("code_verifier", expectedCode),
	)
	if err != nil {
		return nil, NewError(nethttp.StatusInternalServerError, fmt.Sprintf("exchange: %s", err.Error()))
	}

	// verify token
	idToken, err := c.verifier.Verify(r.Context(), token.AccessToken)
	if err != nil {
		return nil, NewError(nethttp.StatusBadRequest, fmt.Sprintf("verify: %s", err.Error()))
	}
	if idToken.Nonce != expectedNonce {
		return nil, NewError(nethttp.StatusBadRequest, "invalid nonce")
	}

	return token, nil
}
