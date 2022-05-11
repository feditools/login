package token

// Kind represents the kind of model to encode a token for
type Kind int64

// This order can not change else all external urls with tokens will become invalid
const (
	// KindFediInstance is a token that represents a federated social instance
	KindFediInstance Kind = 1 + iota
	// KindFediAccount is a token that represents a federated social account
	KindFediAccount
	// KindOauthClient is a token that represents an oauth client
	KindOauthClient
	// KindOauthScope is a token that represents an oauth scope
	KindOauthScope
	// KindApplicationToken is a token that represents an application token
	KindApplicationToken
)

func (k Kind) String() string {
	switch k {
	case KindApplicationToken:
		return "ApplicationToken"
	case KindFediAccount:
		return "FediAccount"
	case KindFediInstance:
		return "FediInstance"
	case KindOauthClient:
		return "OauthClient"
	case KindOauthScope:
		return "OauthScope"
	default:
		return "unknown"
	}
}
