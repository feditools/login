package wellknown

import (
	"encoding/json"
	"github.com/feditools/login/internal/config"
	"github.com/spf13/viper"
	nethttp "net/http"
	"strings"

	"github.com/feditools/go-lib/http"
	"github.com/feditools/login/internal/http/wellknown/models"
	"github.com/feditools/login/internal/path"
)

// OpenidConfigurationGetHandler logs a user out.
func (m *Module) OpenidConfigurationGetHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	l := logger.WithField("func", "OpenidConfigurationGetHandler")

	w.Header().Set(http.HeaderContentType, http.MimeAppJSON)
	cacheControl := []string{
		http.CacheControlNoCache,
		http.CacheControlMustRevalidate,
		http.CacheControlNoTransform,
		http.CacheControlNoStore,
	}
	w.Header().Set(http.HeaderCacheControl, strings.Join(cacheControl, ", "))
	_, err := w.Write(m.openidConfigurationBody)
	if err != nil {
		l.Errorf("writing response: %s", err.Error())
	}
}

func (m *Module) generateOpenidConfigurationBody() error {
	externalURL := strings.TrimSuffix(viper.GetString(config.Keys.ServerExternalURL), "/")

	response := models.OpenidConfiguration{
		Issuer:                externalURL,
		AuthorizationEndpoint: externalURL + path.OauthAuthorize,
		TokenEndpoint:         externalURL + path.OauthToken,
		JwksURI:               externalURL + path.WellKnownOpenidConfigurationJWKS,
	}

	b, err := json.Marshal(response)
	if err != nil {
		return err
	}

	m.openidConfigurationBody = b
	return nil
}
