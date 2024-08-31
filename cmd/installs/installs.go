package installs

import (
	"github.com/robhenley/go-wpe-cli/cmd/installs/domains"
	"github.com/spf13/cobra"
)

// installsCmd represents the installs command
var InstallsCmd = &cobra.Command{
	Use:   "installs",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	InstallsCmd.AddCommand(getCmd)
	InstallsCmd.AddCommand(listCmd)
	InstallsCmd.AddCommand(createCmd)
	InstallsCmd.AddCommand(updateCmd)
	InstallsCmd.AddCommand(deleteCmd)
	InstallsCmd.AddCommand(domains.DomainsCmd)
}
