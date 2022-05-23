package models

// OpenidConfiguration is returned as an oidc response.
type OpenidConfiguration struct {
	Issuer                string   `json:"issuer,omitempty"`
	AuthorizationEndpoint string   `json:"authorization_endpoint,omitempty"`
	TokenEndpoint         string   `json:"token_endpoint,omitempty"`
	JwksURI               string   `json:"jwks_uri,omitempty"`
	TokenSupportedAlgos   []string `json:"id_token_signing_alg_values_supported,omitempty"`
}
