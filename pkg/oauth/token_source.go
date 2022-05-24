package oauth

import (
	"context"

	"golang.org/x/oauth2"
)

// TokenSource renews the token if needed and returns it.
func (c *Client) TokenSource(ctx context.Context, token *oauth2.Token) (*oauth2.Token, bool, error) {
	tokenSource := c.config.TokenSource(ctx, token)

	newToken, err := tokenSource.Token()
	if err != nil {
		return nil, false, err
	}

	return newToken, newToken.AccessToken != token.AccessToken, nil
}
