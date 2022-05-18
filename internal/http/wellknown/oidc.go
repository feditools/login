package wellknown

import (
	"encoding/json"
	nethttp "net/http"
	"strings"

	"github.com/feditools/go-lib/http"
	"github.com/feditools/login/internal/http/wellknown/models"
	"github.com/feditools/login/internal/path"
)

// OpenidConfigurationGetHandler logs a user out.
func (m *Module) OpenidConfigurationGetHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	l := logger.WithField("func", "OpenidConfigurationGetHandler")

	response := models.OpenidConfiguration{
		Issuer:                m.externalURL,
		AuthorizationEndpoint: m.externalURL + path.OauthAuthorize,
		TokenEndpoint:         m.externalURL + path.OauthAuthorize,
	}

	w.Header().Set(http.HeaderContentType, http.MimeAppJSON)
	cacheControl := []string{
		http.CacheControlNoCache,
		http.CacheControlMustRevalidate,
		http.CacheControlNoTransform,
		http.CacheControlNoStore,
	}
	w.Header().Set(http.HeaderCacheControl, strings.Join(cacheControl, ", "))
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		l.Errorf("writing response: %s", err.Error())
	}
}
