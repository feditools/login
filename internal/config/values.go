package config

import "time"

// Values contains the type of each value.
type Values struct {
	ConfigPath string
	LogLevel   string

	// application
	ApplicationName    string
	ApplicationWebsite string
	SoftwareVersion    string
	TokenSalt          string

	// database
	DBType          string
	DBAddress       string
	DBPort          int
	DBUser          string
	DBPassword      string
	DBDatabase      string
	DBTLSMode       string
	DBTLSCACert     string
	DBLoadTestData  bool
	DBEncryptionKey string

	// redis
	RedisAddress  string
	RedisDB       int
	RedisPassword string

	// account
	AccountAccount  string
	AccountAddGroup []string

	// auth
	AccessExpiration  time.Duration
	AccessSecret      string
	RefreshExpiration time.Duration
	RefreshSecret     string

	// server
	ServerExternalHostname string
	ServerGRPCBind         string
	ServerHTTPBind         string
	ServerMinifyHTML       bool
	ServerRoles            []string

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

// Defaults contains the default values.
var Defaults = Values{
	ConfigPath: "",
	LogLevel:   "info",

	// application
	ApplicationName:    "feditools-login",
	ApplicationWebsite: "https://github.com/feditools/login",

	// database
	DBType:         "postgres",
	DBAddress:      "",
	DBPort:         5432,
	DBUser:         "",
	DBPassword:     "",
	DBDatabase:     "ft_login",
	DBTLSMode:      "disable",
	DBTLSCACert:    "",
	DBLoadTestData: false,

	// redis
	RedisAddress:  "localhost:6379",
	RedisDB:       0,
	RedisPassword: "",

	// auth
	AccessExpiration:  time.Minute * 15,
	RefreshExpiration: time.Hour * 24 * 7,

	// server
	ServerExternalHostname: "localhost",
	ServerGRPCBind:         ":50051",
	ServerHTTPBind:         ":5000",
	ServerMinifyHTML:       true,
	ServerRoles: []string{
		ServerRoleWebapp,
	},

	// webapp
	WebappBootstrapCSSURI:         "https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css",
	WebappBootstrapCSSIntegrity:   "sha384-1BmE4kWBq78iYhFldvKuhfTAU6auU8tT94WrHftjDBrCEXSU1oBoqyl2QvZ6jIW3",
	WebappBootstrapJSURI:          "https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/js/bootstrap.bundle.min.js",
	WebappBootstrapJSIntegrity:    "sha384-ka7Sk0Gln4gmtz2MlQnikT1wXgYsOg+OMhuP+IlRH9sENBO0LRn5q+8nbTov4+1p",
	WebappFontAwesomeCSSURI:       "https://cdn.fedi.tools/vendor/fontawesome-free-6.1.1/css/all.min.css",
	WebappFontAwesomeCSSIntegrity: "sha384-/frq1SRXYH/bSyou/HUp/hib7RVN1TawQYja658FEOodR/FQBKVqT9Ol+Oz3Olq5",
	WebappLogoSrcDark:             "https://cdn.fedi.tools/img/feditools-logo-dark.svg",
	WebappLogoSrcLight:            "https://cdn.fedi.tools/img/feditools-logo-light.svg",

	// metrics
	MetricsStatsDAddress: "localhost:8125",
	MetricsStatsDPrefix:  "login",
}
