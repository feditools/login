package flag

import (
	"github.com/feditools/login/internal/config"
	"github.com/spf13/cobra"
)

// Account adds all flags for running the account command.
func Account(cmd *cobra.Command, values config.Values) {
	cmd.PersistentFlags().String(config.Keys.AccountAccount, values.AccountAccount, usage.AccountAccount)
}

// AccountModify adds all flags for running the account modify command.
func AccountModify(cmd *cobra.Command, values config.Values) {
	cmd.PersistentFlags().StringArray(config.Keys.AccountAddGroup, values.AccountAddGroup, usage.AccountAddGroup)
}
