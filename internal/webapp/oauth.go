package webapp

import nethttp "net/http"

// OauthAuthorizeHandler handles oauth authorization
func (m *Module) OauthAuthorizeHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	err := m.oauth.HandleAuthorizeRequest(w, r)
	if err != nil {
		m.returnErrorPage(w, r, nethttp.StatusBadRequest, err.Error())
	}
}

// OauthTokenHandler handles oauth tokens
func (m *Module) OauthTokenHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	err := m.oauth.HandleAuthorizeRequest(w, r)
	if err != nil {
		m.returnErrorPage(w, r, nethttp.StatusBadRequest, err.Error())
	}
}
