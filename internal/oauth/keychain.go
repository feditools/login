package oauth

import (
	"crypto/ecdsa"
	"crypto/sha1" // #nosec G505
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/feditools/login/internal/config"
	"github.com/spf13/viper"
)

// Keychain holds signing keys for.
type Keychain struct {
	ecdsa    *ecdsa.PrivateKey
	ecdsaKID string
}

// NewKeychain creates a new oauth keychain.
func NewKeychain() (*Keychain, error) {
	// read ecdsa key
	ecdsaKey, err := readECDSA()
	if err != nil {
		return nil, err
	}

	// generate ecdsa key id
	ecdsaKeyID := generateKeyID(strings.Join([]string{ecdsaKey.PublicKey.X.String(), ecdsaKey.PublicKey.Y.String()}, "+"))

	return &Keychain{
		ecdsa:    ecdsaKey,
		ecdsaKID: ecdsaKeyID,
	}, nil
}

// keychain utils

func readECDSA() (*ecdsa.PrivateKey, error) {
	// read private key
	privateKeyFile, err := os.ReadFile(viper.GetString(config.Keys.ECPrivateKey))
	if err != nil {
		return nil, err
	}
	privateKeyBlock, _ := pem.Decode(privateKeyFile)
	if err != nil {
		return nil, err
	}
	privateKeyBlockEncoded := privateKeyBlock.Bytes
	privateKey, err := x509.ParseECPrivateKey(privateKeyBlockEncoded)
	if err != nil {
		return nil, err
	}

	// read public key
	publicKeyFile, err := os.ReadFile(viper.GetString(config.Keys.ECPublicKey))
	if err != nil {
		return nil, err
	}
	publicKeyBlock, _ := pem.Decode(publicKeyFile)
	if err != nil {
		return nil, err
	}
	publicKeyBlockEncoded := publicKeyBlock.Bytes
	publicKeyI, err := x509.ParsePKIXPublicKey(publicKeyBlockEncoded)
	if err != nil {
		return nil, err
	}
	publicKey, ok := publicKeyI.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("public key not ecdsa")
	}

	// join ecdsa keys
	privateKey.PublicKey = *publicKey

	return privateKey, nil
}

func generateKeyID(s string) string {
	h := sha1.New() // #nosec G401 not used for cryptography
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
