package main

import (
	"github.com/feditools/login/cmd/login/action/server"
	"github.com/feditools/login/cmd/login/flag"
	"github.com/feditools/login/internal/config"
	"github.com/spf13/cobra"
)

// serverCommands returns the 'server' subcommand.
func serverCommands() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "run a feditools server",
	}

	serverStartCmd := &cobra.Command{
		Use:   "start",
		Short: "start the feditools server",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return preRun(cmd)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd.Context(), server.Start)
		},
	}

	flag.Server(serverStartCmd, config.Defaults)

	serverCmd.AddCommand(serverStartCmd)

	return serverCmd
}
