package config

// KeyNames is a struct that contains the names of keys.
type KeyNames struct {
	LogLevel   string
	ConfigPath string

	// application
	ApplicationName    string
	ApplicationWebsite string
	SoftwareVersion    string
	TokenSalt          string

	// database
	DBType          string
	DBAddress       string
	DBPort          string
	DBUser          string
	DBPassword      string
	DBDatabase      string
	DBTLSMode       string
	DBTLSCACert     string
	DBLoadTestData  string
	DBEncryptionKey string

	// redis
	RedisAddress  string
	RedisDB       string
	RedisPassword string

	// account
	AccountAccount  string
	AccountAddGroup string

	// auth
	AccessExpiration  string
	AccessSecret      string
	RefreshExpiration string
	RefreshSecret     string

	// server
	ServerExternalURL string
	ServerGRPCBind    string
	ServerHTTPBind    string
	ServerMinifyHTML  string
	ServerRoles       string

	// webapp
	WebappBootstrapCSSURI         string
	WebappBootstrapCSSIntegrity   string
	WebappBootstrapJSURI          string
	WebappBootstrapJSIntegrity    string
	WebappFontAwesomeCSSURI       string
	WebappFontAwesomeCSSIntegrity string
	WebappLogoSrcDark             string
	WebappLogoSrcLight            string

	// metrics
	MetricsStatsDAddress string
	MetricsStatsDPrefix  string
}

// Keys contains the names of config keys.
var Keys = KeyNames{
	ConfigPath: "config-path", // CLI only
	LogLevel:   "log-level",

	// application
	ApplicationName:    "application-name",
	ApplicationWebsite: "application-website",
	SoftwareVersion:    "software-version", // Set at build
	TokenSalt:          "token-salt",

	// database
	DBType:          "db-type",
	DBAddress:       "db-address",
	DBPort:          "db-port",
	DBUser:          "db-user",
	DBPassword:      "db-password",
	DBDatabase:      "db-database",
	DBTLSMode:       "db-tls-mode",
	DBTLSCACert:     "db-tls-ca-cert",
	DBLoadTestData:  "test-data", // CLI only
	DBEncryptionKey: "db-crypto-key",

	// redis
	RedisAddress:  "redis-address",
	RedisDB:       "redis-db",
	RedisPassword: "redis-password",

	// account
	AccountAccount:  "account",
	AccountAddGroup: "add-group",

	// auth
	AccessExpiration:  "access-expiration",
	AccessSecret:      "access-secret",
	RefreshExpiration: "refresh-expiration",
	RefreshSecret:     "refresh-secret",

	// server
	ServerExternalURL: "external-url",
	ServerGRPCBind:    "grpc-bind",
	ServerHTTPBind:    "http-bind",
	ServerMinifyHTML:  "minify-html",
	ServerRoles:       "server-role",

	// webapp
	WebappBootstrapCSSURI:         "webapp-bootstrap-css-uri",
	WebappBootstrapCSSIntegrity:   "webapp-bootstrap-css-integrity",
	WebappBootstrapJSURI:          "webapp-bootstrap-js-uri",
	WebappBootstrapJSIntegrity:    "webapp-bootstrap-js-integrity",
	WebappFontAwesomeCSSURI:       "webapp-fontawesome-css-uri",
	WebappFontAwesomeCSSIntegrity: "webapp-fontawesome-css-integrity",
	WebappLogoSrcDark:             "webapp-logo-src-dark",
	WebappLogoSrcLight:            "webapp-logo-src-light",

	// metrics
	MetricsStatsDAddress: "statsd-addr",
	MetricsStatsDPrefix:  "statsd-prefix",
}
