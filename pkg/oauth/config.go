package oauth

// Config contains configuration for the oauth client.
type Config struct {
	CallbackURL  string
	ServerURL    string
	ClientID     string
	ClientSecret string
}
