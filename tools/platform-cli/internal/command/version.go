package command

import (
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:  "version",
	Long: "Print the version number of platform-cli",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("platform-cli v0.1.0")
	},
}
