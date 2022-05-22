package wellknown

import (
	"encoding/json"
	nethttp "net/http"
	"strings"

	"github.com/feditools/go-lib/http"
	jose "gopkg.in/square/go-jose.v2"
)

// OpenidConfigurationJWKSGetHandler logs a user out.
func (m *Module) OpenidConfigurationJWKSGetHandler(w nethttp.ResponseWriter, _ *nethttp.Request) {
	l := logger.WithField("func", "OpenidConfigurationJWKSGetHandler")

	w.Header().Set(http.HeaderContentType, http.MimeAppJSON)
	cacheControl := []string{
		http.CacheControlNoCache,
	}
	w.Header().Set(http.HeaderCacheControl, strings.Join(cacheControl, ", "))
	_, err := w.Write(m.openidConfigurationJWKSBody)
	if err != nil {
		l.Errorf("writing response: %s", err.Error())
	}
}

func (m *Module) generateOpenidConfigurationJWKSBody() error {
	l := logger.WithField("func", "generateOpenidConfigurationJWKSBody")

	var keys []jose.JSONWebKey

	// generate ec public key
	ecKey, err := m.generateECPublicKeyJSONWebKey()
	if err != nil {
		l.Errorf("generate ec public key: %s", err.Error())

		return err
	}
	keys = append(keys, *ecKey)

	// create response
	response := jose.JSONWebKeySet{
		Keys: keys,
	}

	b, err := json.Marshal(response)
	if err != nil {
		l.Errorf("json marshal: %s", err.Error())

		return err
	}

	m.openidConfigurationJWKSBody = b

	return nil
}

func (m *Module) generateECPublicKeyJSONWebKey() (*jose.JSONWebKey, error) {
	// l := logger.WithField("func", "generateECPublicKeyJSONWebKey")

	publicKey := m.oauth.GetECPublicKey()

	return &jose.JSONWebKey{
		Key:   publicKey,
		KeyID: m.oauth.GetECPublicKeyID(),
	}, nil
}
