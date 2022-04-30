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
	cmd.PersistentFlags().String(config.Keys.ServerHTTPBind, values.ServerHTTPBind, usage.ServerHTTPBind)
	cmd.PersistentFlags().Bool(config.Keys.ServerMinifyHTML, values.ServerMinifyHTML, usage.ServerMinifyHTML)
	cmd.PersistentFlags().StringArray(config.Keys.ServerRoles, values.ServerRoles, usage.ServerRoles)

	// auth
	cmd.PersistentFlags().Duration(config.Keys.AccessExpiration, values.AccessExpiration, usage.AccessExpiration)
	cmd.PersistentFlags().String(config.Keys.AccessSecret, values.AccessSecret, usage.AccessSecret)
	cmd.PersistentFlags().Duration(config.Keys.RefreshExpiration, values.RefreshExpiration, usage.RefreshExpiration)
	cmd.PersistentFlags().String(config.Keys.RefreshSecret, values.RefreshSecret, usage.RefreshSecret)
}
