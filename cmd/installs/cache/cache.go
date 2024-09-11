package cache

import (
	"github.com/spf13/cobra"
)

// InstallsCacheCmd represents the create command
var InstallsCacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Commands related to install caching",
	Long:  `Commands related to install caching`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func init() {
	InstallsCacheCmd.AddCommand(installsCachePurgeCmd)
}
