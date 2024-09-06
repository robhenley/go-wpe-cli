package backups

import (
	"github.com/spf13/cobra"
)

// BackupsCmd represents the accounts command
var BackupsCmd = &cobra.Command{
	Use:   "backups",
	Short: "Commands related to backing up an install",
	Long:  `Use this as an entry point to creating backups of your installs.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Usage()
			return
		}
	},
}

func init() {
	BackupsCmd.AddCommand(installsBackupsCreateCmd)
	BackupsCmd.AddCommand(installsBackupsGetCmd)
}
