package oauth

import (
	"errors"
	nethttp "net/http"
	"strconv"
)

// HandleAuthorizeRequest passes an authorize request to the oauth server.
func (s *Server) HandleAuthorizeRequest(uid int64, w nethttp.ResponseWriter, r *nethttp.Request) error {
	l := logger.WithField("func", "HandleAuthorizeRequest")

	nonce := r.Form.Get("nonce")
	if nonce == "" {
		return errors.New("missing nonce")
	}

	err := s.kv.SetOauthNonceLogin(r.Context(), strconv.FormatInt(uid, 10), nonce, s.LoginNonceExp)
	if err != nil {
		l.Errorf("set oauth nonce: %s", err.Error())

		return err
	}

	return s.oauth.HandleAuthorizeRequest(w, r)
}
