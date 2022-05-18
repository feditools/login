package wellknown

import (
	"context"
	"strings"

	"github.com/feditools/login/internal/config"
	"github.com/feditools/login/internal/http"
	"github.com/spf13/viper"
)

// Module contains a well-known module for the web server. Implements web.Module.
type Module struct {
	srv *http.Server

	externalURL string
}

// New creates a new.
func New(ctx context.Context) (*Module, error) {
	return &Module{
		externalURL: strings.TrimSuffix(viper.GetString(config.Keys.ServerExternalURL), "/"),
	}, nil
}

// Name return the module name.
func (*Module) Name() string {
	return config.ServerRoleWellKnown
}

// SetServer adds a reference to the server to the module.
func (m *Module) SetServer(s *http.Server) {
	m.srv = s
}
