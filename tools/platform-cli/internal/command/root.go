package command

import (
	"github.com/rsturla/platform/tools/platform-cli/pkg/log"
	"github.com/spf13/cobra"
)

var logLevel string

var rootCmd = &cobra.Command{
	Use:   "platform-cli",
	Short: "A CLI for interacting with the Platform monorepo and its services",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return log.ConfigureLogger(logLevel)
	},
}
