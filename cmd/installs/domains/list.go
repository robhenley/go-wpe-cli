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

// installsDomainsListCmd represents the domains list command
var installsDomainsListCmd = &cobra.Command{
	Use:   "list <install id>",
	Short: "Get the domains for an install by install ID",
	Long:  `Returns domains for a specific install.`,
	Run:   installsDomainsList,
}

func init() {
	installsDomainsListCmd.Flags().Int("page", 1, "The page to return")
	installsDomainsListCmd.Flags().Int("limit", 100, "Limit the number of results")
}

func installsDomainsList(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		cmd.Usage()
		return
	}

	installID := args[0]

	page, err := cmd.Flags().GetInt("page")
	cobra.CheckErr(err)

	limit, err := cmd.Flags().GetInt("limit")
	cobra.CheckErr(err)

	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	config.Limit = limit

	api := api.NewAPI(config)
	domains, err := api.InstallDomainsList(installID, page)
	cobra.CheckErr(err)

	format, err := cmd.Flags().GetString("format")
	cobra.CheckErr(err)

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(domains)
		cobra.CheckErr(err)

		fmt.Fprintf(os.Stdout, "%s", j)
		return
	}

	if len(domains) == 0 {
		cmd.PrintErr("Error: No domains were returned")
	}

	for _, domain := range domains {
		fmt.Fprintf(os.Stdout, "%s %-40s\n", domain.ID, domain.Name)
	}

}
