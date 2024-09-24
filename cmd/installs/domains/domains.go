package domains

import (
	"github.com/spf13/cobra"
)

// installsDomainsCmd represents the installsDomains command
var DomainsCmd = &cobra.Command{
	Use:   "domains",
	Short: "Manage your installs domains",
	Long:  `Manage your installs domains`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	DomainsCmd.AddCommand(installsDomainsGetCmd)
	DomainsCmd.AddCommand(installsDomainsDeleteCmd)
	DomainsCmd.AddCommand(installsDomainsListCmd)
	DomainsCmd.AddCommand(installsDomainsCreateCmd)
	DomainsCmd.AddCommand(installsDomainsUpdateCmd)
	DomainsCmd.AddCommand(installsDomainsCdnStatusCmd)
	DomainsCmd.AddCommand(installsDomainsBulkCreateCmd)
}
