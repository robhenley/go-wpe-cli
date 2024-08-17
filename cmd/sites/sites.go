/*
Copyright © 2024 Rob Henley <rob.henley@gmail.com>
*/
package sites

import (
	"os"

	"github.com/spf13/cobra"
)

// sitesCmd represents the sites command
var SitesCmd = &cobra.Command{
	Use:   "sites",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
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
