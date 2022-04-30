package webapp

import (
	"github.com/feditools/login"
	"github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/path"
	iofs "io/fs"
	nethttp "net/http"
)

// Route attaches routes to the web server
func (m *Module) Route(s *http.Server) error {
	staticFS, err := iofs.Sub(login.Files, DirStatic)
	if err != nil {
		return err
	}

	// Static Files
	s.PathPrefix(path.Static).Handler(nethttp.StripPrefix(path.Static, nethttp.FileServer(nethttp.FS(staticFS))))

	webapp := s.PathPrefix("/").Subrouter()
	webapp.Use(m.Middleware)
	webapp.NotFoundHandler = m.notFoundHandler()
	webapp.MethodNotAllowedHandler = m.methodNotAllowedHandler()

	webapp.HandleFunc(path.CallbackOauth, m.CallbackOauthGetHandler).Methods("GET")
	webapp.HandleFunc(path.Login, m.LoginGetHandler).Methods("GET")
	webapp.HandleFunc(path.OauthAuthorize, m.OauthAuthorizeHandler)
	webapp.HandleFunc(path.OauthToken, m.OauthTokenHandler)

	return nil
}
