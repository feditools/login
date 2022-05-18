package models

const (
	// KeyTypeEC indicated the key is a ecdsa type.
	KeyTypeEC = "EC"
	// KeyTypeRSA indicated the key is a rsa type.
	KeyTypeRSA = "RSA"

	// UseEncryption indicates the key is used for encryption.
	UseEncryption = "enc"
	// UseSignature indicates the key is used for signatures.
	UseSignature = "sig"
)
