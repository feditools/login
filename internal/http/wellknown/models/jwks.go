package models

// JWKS is returned as an jwks response.
type JWKS struct {
	Keys []JWKSKey `json:"keys"`
}

// JWKSKey is returned as an jwks key.
type JWKSKey struct {
	KeyType       string `json:"kty,omitempty"`
	PublicKeyUse  string `json:"use,omitempty"`
	KeyOperations string `json:"key_ops,omitempty"`
	Algorithm     string `json:"alg,omitempty"`
	KeyID         string `json:"kid,omitempty"`

	// ecdsa
	Curve string `json:"crv,omitempty"`
	X     string `json:"x,omitempty"`
	Y     string `json:"y,omitempty"`

	// rsa
	N string `json:"n,omitempty"`
	E string `json:"e,omitempty"`

	// x509
	X509URL                         string   `json:"x5u,omitempty"`
	X509CertificateChain            []string `json:"x5c,omitempty"`
	X509CertificateThumbprintSHA1   string   `json:"x5t,omitempty"`
	X509CertificateThumbprintSHA256 string   `json:"x5t#S256,omitempty"`
}
