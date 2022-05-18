package wellknown

import (
	nethttp "net/http"

	"github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/path"
)

// Route attaches routes to the web server.
func (m *Module) Route(s *http.Server) error {
	wellknown := s.PathPrefix(path.WellKnown).Subrouter()
	wellknown.NotFoundHandler = m.notFoundHandler()
	wellknown.MethodNotAllowedHandler = m.methodNotAllowedHandler()

	wellknown.HandleFunc(path.WellKnownSubOpenidConfiguration, m.OpenidConfigurationGetHandler).Methods(nethttp.MethodGet)
	wellknown.HandleFunc(path.WellKnownSubOpenidConfigurationJWKS, m.OpenidConfigurationJWKSGetHandler).Methods(nethttp.MethodGet)
	return nil
}
