package token

// Kind represents the kind of model to encode a token for
type Kind int64

const (
	// KindFediInstance is a token that represents a federated social instance
	KindFediInstance Kind = 1 + iota
	// KindFediAccount is a token that represents a federated social account
	KindFediAccount
)

func (k Kind) String() string {
	switch k {
	case KindFediInstance:
		return "FediInstance"
	case KindFediAccount:
		return "FediAccount"
	default:
		return "unknown"
	}
}
