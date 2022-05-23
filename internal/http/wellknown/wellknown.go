package wellknown

import (
	"context"

	"github.com/feditools/login/internal/config"
	"github.com/feditools/login/internal/http"
	"github.com/feditools/login/internal/oauth"
)

// Module contains a well-known module for the web server. Implements web.Module.
type Module struct {
	oauth *oauth.Server
	srv   *http.Server

	openidConfigurationBody     []byte
	openidConfigurationJWKSBody []byte
}

// New creates a new well-known module.
func New(_ context.Context, oauthServer *oauth.Server) (*Module, error) {
	module := &Module{
		oauth: oauthServer,
	}

	err := module.generateOpenidConfigurationBody()
	if err != nil {
		return nil, err
	}
	err = module.generateOpenidConfigurationJWKSBody()
	if err != nil {
		return nil, err
	}

	return module, nil
}

// Name return the module name.
func (*Module) Name() string {
	return config.ServerRoleWellKnown
}

// SetServer adds a reference to the server to the module.
func (m *Module) SetServer(s *http.Server) {
	m.srv = s
}
