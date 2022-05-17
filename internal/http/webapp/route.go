package webapp

import (
	iofs "io/fs"
	nethttp "net/http"

	"github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/path"
	"github.com/feditools/login/web"
)

// Route attaches routes to the web server.
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

	webapp.HandleFunc(path.CallbackOauth, m.CallbackOauthGetHandler).Methods(nethttp.MethodGet)
	webapp.HandleFunc(path.Login, m.LoginGetHandler).Methods(nethttp.MethodGet)
	webapp.HandleFunc(path.Login, m.LoginPostHandler).Methods(nethttp.MethodPost)
	webapp.HandleFunc(path.Logout, m.LogoutGetHandler).Methods(nethttp.MethodGet)
	webapp.HandleFunc(path.OauthAuthorize, m.OauthAuthorizeGetHandler).Methods(nethttp.MethodGet)
	webapp.HandleFunc(path.OauthToken, m.OauthTokenHandler)

	admin := webapp.PathPrefix(path.Admin).Subrouter()
	admin.Use(m.MiddlewareRequireAdmin)
	admin.NotFoundHandler = m.notFoundHandler()
	admin.MethodNotAllowedHandler = m.methodNotAllowedHandler()

	admin.HandleFunc(path.AdminSubFediverse, m.AdminFediGetHandler).Methods(nethttp.MethodGet)
	admin.HandleFunc(path.AdminSubFediverseAccounts, m.AdminFediAccountsGetHandler).Methods(nethttp.MethodGet)
	admin.HandleFunc(path.AdminSubFediverseInstances, m.AdminFediInstancesGetHandler).Methods(nethttp.MethodGet)

	admin.HandleFunc(path.AdminSubOauth, m.AdminOauthGetHandler).Methods(nethttp.MethodGet)
	admin.HandleFunc(path.AdminSubOauthClientAdd, m.AdminOauthClientAddGetHandler).Methods(nethttp.MethodGet)
	admin.HandleFunc(path.AdminSubOauthClientAdd, m.AdminOauthClientAddPostHandler).Methods(nethttp.MethodPost)
	admin.HandleFunc(path.AdminSubOauthClients, m.AdminOauthClientsGetHandler).Methods(nethttp.MethodGet)

	admin.HandleFunc(path.AdminSubSystem, m.AdminSystemGetHandler).Methods(nethttp.MethodGet)
	admin.HandleFunc(path.AdminSubSystemApplicationTokens, m.AdminSystemApplicationTokensGetHandler).Methods(nethttp.MethodGet)
	admin.HandleFunc(path.AdminSubSystemApplicationTokenAdd, m.AdminSystemApplicationTokenAddGetHandler).Methods(nethttp.MethodGet)
	admin.HandleFunc(path.AdminSubSystemApplicationTokenAdd, m.AdminSystemApplicationTokenAddPostHandler).Methods(nethttp.MethodPost)

	return nil
}
