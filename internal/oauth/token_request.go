package oauth

import nethttp "net/http"

// HandleTokenRequest passes a token request to the oauth server.
func (s *Server) HandleTokenRequest(w nethttp.ResponseWriter, r *nethttp.Request) error {
	return s.oauth.HandleTokenRequest(w, r)
}
