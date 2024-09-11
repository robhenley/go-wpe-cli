package installs

import (
	"github.com/robhenley/go-wpe-cli/cmd/installs/backups"
	"github.com/robhenley/go-wpe-cli/cmd/installs/cache"
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
	InstallsCmd.AddCommand(installsGetCmd)
	InstallsCmd.AddCommand(installsListCmd)
	InstallsCmd.AddCommand(installsCreateCmd)
	InstallsCmd.AddCommand(installsUpdateCmd)
	InstallsCmd.AddCommand(installsDeleteCmd)
	InstallsCmd.AddCommand(cache.InstallsCacheCmd)
	InstallsCmd.AddCommand(domains.DomainsCmd)
	InstallsCmd.AddCommand(backups.BackupsCmd)
}
