package cdn

import (
	"github.com/robhenley/go-wpe-cli/cmd/types"
	"github.com/robhenley/go-wpe-cli/internal/api"
	"github.com/spf13/cobra"
)

// cdnStatusCmd represents the CDN status command
var cdnStatusCmd = &cobra.Command{
	Use:   "status <install id> <domain id>",
	Short: "Check the status of a domain",
	Long:  `Submits a request to check the status of the domain`,
	Run:   cdnStatus,
}

func cdnStatus(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		cmd.Usage()
		return
	}

	installID := args[0]
	domainID := args[1]

	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)

	api := api.NewAPI(config)
	res, err := api.InstallDomainCDNStatus(installID, domainID)
	if err != nil {
		cmd.PrintErrln(err)
		return
	}

	cmd.Println(res)

}
