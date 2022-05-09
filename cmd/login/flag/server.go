package flag

import (
	"github.com/feditools/login/internal/config"
	"github.com/spf13/cobra"
)

// Server adds all flags for running the server.
func Server(cmd *cobra.Command, values config.Values) {
	Redis(cmd, values)

	// server
	cmd.PersistentFlags().String(config.Keys.ServerExternalHostname, values.ServerExternalHostname, usage.ServerExternalHostname)
	cmd.PersistentFlags().String(config.Keys.ServerGRPCBind, values.ServerGRPCBind, usage.ServerGRPCBind)
	cmd.PersistentFlags().String(config.Keys.ServerHTTPBind, values.ServerHTTPBind, usage.ServerHTTPBind)
	cmd.PersistentFlags().Bool(config.Keys.ServerMinifyHTML, values.ServerMinifyHTML, usage.ServerMinifyHTML)
	cmd.PersistentFlags().StringArray(config.Keys.ServerRoles, values.ServerRoles, usage.ServerRoles)

	// auth
	cmd.PersistentFlags().Duration(config.Keys.AccessExpiration, values.AccessExpiration, usage.AccessExpiration)
	cmd.PersistentFlags().String(config.Keys.AccessSecret, values.AccessSecret, usage.AccessSecret)
	cmd.PersistentFlags().Duration(config.Keys.RefreshExpiration, values.RefreshExpiration, usage.RefreshExpiration)
	cmd.PersistentFlags().String(config.Keys.RefreshSecret, values.RefreshSecret, usage.RefreshSecret)

	// webapp
	cmd.PersistentFlags().String(config.Keys.WebappBootstrapCSSURI, values.WebappBootstrapCSSURI, usage.WebappBootstrapCSSURI)
	cmd.PersistentFlags().String(config.Keys.WebappBootstrapCSSIntegrity, values.WebappBootstrapCSSIntegrity, usage.WebappBootstrapCSSIntegrity)
	cmd.PersistentFlags().String(config.Keys.WebappBootstrapJSURI, values.WebappBootstrapJSURI, usage.WebappBootstrapJSURI)
	cmd.PersistentFlags().String(config.Keys.WebappBootstrapJSIntegrity, values.WebappBootstrapJSIntegrity, usage.WebappBootstrapJSIntegrity)
	cmd.PersistentFlags().String(config.Keys.WebappFontAwesomeCSSURI, values.WebappFontAwesomeCSSURI, usage.WebappFontAwesomeCSSURI)
	cmd.PersistentFlags().String(config.Keys.WebappFontAwesomeCSSIntegrity, values.WebappFontAwesomeCSSIntegrity, usage.WebappFontAwesomeCSSIntegrity)
	cmd.PersistentFlags().String(config.Keys.WebappLogoSrcDark, values.WebappLogoSrcDark, usage.WebappLogoSrcDark)
	cmd.PersistentFlags().String(config.Keys.WebappLogoSrcLight, values.WebappLogoSrcLight, usage.WebappLogoSrcLight)
}
