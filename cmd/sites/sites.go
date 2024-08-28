/*
Copyright © 2024 Rob Henley <rob.henley@gmail.com>
*/
package sites

import (
	"github.com/spf13/cobra"
)

// sitesCmd represents the sites command
var SitesCmd = &cobra.Command{
	Use:   "sites",
	Short: "All things site related (e.g. retrieval, listing, creation, deletion, etc)",
	Long:  `Long.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
	},
}

func init() {
	SitesCmd.AddCommand(listCmd)
	SitesCmd.AddCommand(getCmd)
	SitesCmd.AddCommand(createCmd)
	SitesCmd.AddCommand(updateCmd)
	SitesCmd.AddCommand(deleteCmd)
}
