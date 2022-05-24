package oauth

import (
	"context"
	"fmt"
	nethttp "net/http"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gorilla/sessions"

	"golang.org/x/oauth2"
)

// TokenSource renews the token if needed and returns it.
func (c *Client) TokenSource(ctx context.Context, us *sessions.Session, token *oauth2.Token) (*oauth2.Token, *oidc.IDToken, bool, error) {
	tokenSource := c.config.TokenSource(ctx, token)

	// get possibly renewed token
	newToken, err := tokenSource.Token()
	if err != nil {
		return nil, nil, false, NewError(nethttp.StatusInternalServerError, err.Error())
	}

	// verify token if new
	if newToken.AccessToken != token.AccessToken {
		expectedNonce, ok := us.Values[SessionKeyNonce].(string)
		if !ok {
			return nil, nil, false, NewError(nethttp.StatusBadRequest, "missing nonce")
		}

		idToken, err := c.verifier.Verify(ctx, token.AccessToken)
		if err != nil {
			return nil, nil, false, NewError(nethttp.StatusBadRequest, fmt.Sprintf("verify: %s", err.Error()))
		}
		if idToken.Nonce != expectedNonce {
			return nil, nil, false, NewError(nethttp.StatusBadRequest, "invalid nonce")
		}
	}
	return nil, nil, false, nil
}
