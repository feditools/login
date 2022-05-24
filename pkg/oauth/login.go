package oauth

import (
	"fmt"
	nethttp "net/http"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
)

// HandleLogin sends the user to the oauth login server.
func (c *Client) HandleLogin(w nethttp.ResponseWriter, r *nethttp.Request, us *sessions.Session) error {
	newCode := uuid.New().String()
	newNonce := uuid.New().String()
	newState := uuid.New().String()
	us.Values[SessionKeyCode] = newCode
	us.Values[SessionKeyNonce] = newNonce
	us.Values[SessionKeyState] = newState
	if err := us.Save(r, w); err != nil {
		return NewError(nethttp.StatusInternalServerError, fmt.Sprintf("session: %s", err.Error()))
	}

	authCodeURL := c.config.AuthCodeURL(
		newState,
		oidc.Nonce(newNonce),
		oauth2.SetAuthURLParam("code_challenge", genCodeChallengeS256(newCode)),
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
	)
	nethttp.Redirect(w, r, authCodeURL, nethttp.StatusFound)

	return nil
}
