package domains

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/robhenley/go-wpe-cli/cmd/types"
	"github.com/robhenley/go-wpe-cli/internal/api"
	"github.com/spf13/cobra"
)

// installsDomainsGetCmd represents the installsDomainsGet command
var installsDomainsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a specific domain for a given install",
	Long:  `Returns a specific domain for a given install`,
	Run:   installsDomainsGet,
}

func init() {
	installsDomainsGetCmd.Flags().StringP("install", "i", "", "The install ID to create a backup from")
	installsDomainsGetCmd.MarkFlagRequired("install")

	installsDomainsGetCmd.Flags().StringP("domain", "d", "", "The domain ID to check the status of")
	installsDomainsGetCmd.MarkFlagRequired("domain")
}

func installsDomainsGet(cmd *cobra.Command, args []string) {
	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)

	installID, err := cmd.Flags().GetString("install")
	cobra.CheckErr(err)

	domainID, err := cmd.Flags().GetString("domain")
	cobra.CheckErr(err)

	format, err := cmd.Flags().GetString("format")
	cobra.CheckErr(err)

	domain, err := api.InstallsDomainsGet(installID, domainID)
	cobra.CheckErr(err)

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(domain)
		cobra.CheckErr(err)

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return
	}

	fmt.Fprintf(os.Stdout, "%s %s %t %s", domain.ID, domain.Name, domain.Primary, domain.RedirectTo.Name)
}
