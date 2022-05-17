package flag

import (
	"github.com/feditools/login/internal/config"
	"github.com/spf13/cobra"
)

// Database adds all flags for running the database command.
func Database(_ *cobra.Command, _ config.Values) {
}

// DatabaseMigrate adds all flags for running the database migrate command.
func DatabaseMigrate(cmd *cobra.Command, values config.Values) {
	Database(cmd, values)
	cmd.PersistentFlags().Bool(config.Keys.DBLoadTestData, values.DBLoadTestData, usage.DBLoadTestData)
}
