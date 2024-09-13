package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "goreleaser"
)

// versionCmd represents the version of the this command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "The version of this command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s (%s) %s %s\n", version, commit, date, builtBy)
	},
}
