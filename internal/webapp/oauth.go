package webapp

import (
	nethttp "net/http"
)

// OauthAuthorizeGetHandler handles oauth authorization
func (m *Module) OauthAuthorizeGetHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	_, shouldReturn := m.authRequireLoggedIn(w, r)
	if shouldReturn {
		return
	}

	err := m.oauth.HandleAuthorizeRequest(w, r)
	if err != nil {
		m.returnErrorPage(w, r, nethttp.StatusBadRequest, err.Error())
	}
}

// OauthTokenHandler handles oauth tokens
func (m *Module) OauthTokenHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	m.oauth.HandleTokenRequest(w, r)
}
