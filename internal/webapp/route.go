package webapp

import (
	"github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/path"
	"github.com/feditools/login/web"
	iofs "io/fs"
	nethttp "net/http"
)

// Route attaches routes to the web server
func (m *Module) Route(s *http.Server) error {
	staticFS, err := iofs.Sub(web.Files, DirStatic)
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
	webapp.HandleFunc(path.Login, m.LoginPostHandler).Methods("POST")
	webapp.HandleFunc(path.Logout, m.LogoutGetHandler).Methods("GET")
	webapp.HandleFunc(path.OauthAuthorize, m.OauthAuthorizeGetHandler).Methods("GET")
	webapp.HandleFunc(path.OauthToken, m.OauthTokenHandler)

	admin := webapp.PathPrefix(path.Admin).Subrouter()
	admin.Use(m.MiddlewareRequireAdmin)
	admin.NotFoundHandler = m.notFoundHandler()
	admin.MethodNotAllowedHandler = m.methodNotAllowedHandler()
	admin.HandleFunc(path.AdminSubOauthClientAdd, m.AdminOauthClientAddGetHandler).Methods("GET")
	admin.HandleFunc(path.AdminSubOauthClientAdd, m.AdminOauthClientAddPostHandler).Methods("POST")
	admin.HandleFunc(path.AdminSubOauthClients, m.AdminOauthClientsGetHandler).Methods("GET")

	return nil
}
