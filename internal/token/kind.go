package token

// Kind represents the kind of model to encode a token for
type Kind int64

const (
	// KindFediInstance is a token that represents a federated social instance
	KindFediInstance Kind = 1 + iota
	// KindFediAccount is a token that represents a federated social account
	KindFediAccount
	// KindOauthClient is a token that represents an oauth client
	KindOauthClient
)

func (k Kind) String() string {
	switch k {
	case KindFediInstance:
		return "FediInstance"
	case KindFediAccount:
		return "FediAccount"
	case KindOauthClient:
		return "OauthClient"
	default:
		return "unknown"
	}
}
