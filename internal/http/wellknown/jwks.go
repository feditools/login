package wellknown

import (
	"crypto/sha1" // #nosec G505
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/base64"
	"encoding/json"
	nethttp "net/http"
	"strings"

	"github.com/feditools/go-lib/http"
	"github.com/feditools/login/internal/http/wellknown/models"
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

	var keys []models.JWKSKey

	// generate ec public key
	ecKey, err := m.generateECPublicKeyJWKSKey()
	if err != nil {
		l.Errorf("generate ec public key: %s", err.Error())
		return err
	}
	keys = append(keys, *ecKey)

	// create response
	response := models.JWKS{
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

func (m *Module) generateECPublicKeyJWKSKey() (*models.JWKSKey, error) {
	l := logger.WithField("func", "generateECPublicKeyJWKSKey")

	publicKey := m.oauth.GetECPublicKey()
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		l.Errorf("MarshalPKIXPublicKey: %s", err.Error())
		return nil, err
	}

	// encode to der
	publicKeyDER, err := asn1.Marshal(publicKeyBytes)
	if err != nil {
		l.Errorf("marshal asn1: %s", err.Error())
		return nil, err
	}

	publicKeyB64 := base64.StdEncoding.EncodeToString(publicKeyBytes)
	publicKeyXB64 := base64.StdEncoding.EncodeToString(m.oauth.GetECPublicKeyX())
	publicKeyYB64 := base64.StdEncoding.EncodeToString(m.oauth.GetECPublicKeyY())

	return &models.JWKSKey{
		KeyType:      models.KeyTypeEC,
		PublicKeyUse: models.UseSignature,
		KeyID:        m.oauth.GetECPublicKeyID(),

		Curve: m.oauth.GetECPublicKeyCurve(),
		X:     publicKeyXB64,
		Y:     publicKeyYB64,

		X509CertificateChain: []string{
			publicKeyB64,
		},
		X509CertificateThumbprintSHA1:   generateX509ThumbprintSHA1(publicKeyDER),
		X509CertificateThumbprintSHA256: generateX509ThumbprintSHA256(publicKeyDER),
	}, nil
}

func generateX509ThumbprintSHA1(publicKeyDER []byte) string {
	// make hash
	hasher := sha1.New() // #nosec G401 not used for cryptography
	hasher.Write(publicKeyDER)
	publicKeySignature := hasher.Sum(nil)

	// base64 encode
	publicKeyB64 := base64.StdEncoding.EncodeToString(publicKeySignature)

	return publicKeyB64
}

func generateX509ThumbprintSHA256(publicKeyDER []byte) string {
	// make hash
	hasher := sha256.New()
	hasher.Write(publicKeyDER)
	publicKeySignature := hasher.Sum(nil)

	// base64 encode
	publicKeyB64 := base64.StdEncoding.EncodeToString(publicKeySignature)

	return publicKeyB64
}
