package domains

import (
	"fmt"
	"os"

	"github.com/robhenley/go-wpe-cli/cmd/types"
	"github.com/robhenley/go-wpe-cli/internal/api"
	"github.com/spf13/cobra"
)

// installsDomainsCdnStatusCmd represents the cdn command
var installsDomainsCdnStatusCmd = &cobra.Command{
	Use:   "cdn-status",
	Short: "Check the status of a domain",
	Long:  `Submits a request to check the status of the domain`,
	Run:   installsDomainsCdnStatus,
}

func init() {
	installsDomainsCdnStatusCmd.Flags().StringP("install", "i", "", "The install ID to create a backup from")
	installsDomainsCdnStatusCmd.MarkFlagRequired("install")

	installsDomainsCdnStatusCmd.Flags().StringP("domain", "d", "", "The domain ID to check the status of")
	installsDomainsCdnStatusCmd.MarkFlagRequired("domain")
}

func installsDomainsCdnStatus(cmd *cobra.Command, args []string) {
	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)

	installID, err := cmd.Flags().GetString("install")
	cobra.CheckErr(err)

	domainID, err := cmd.Flags().GetString("domain")
	cobra.CheckErr(err)

	res, err := api.InstallDomainCDNStatus(installID, domainID)
	cobra.CheckErr(err)

	// TODO: Review this
	cmd.Println(res)

	fmt.Fprintf(os.Stdout, "%s, %s", installID, domainID)

}
