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

// installsDomainsDeleteCmd represents the installsDomainsDelete command
var installsDomainsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a specific domain for a given install",
	Long:  `Delete a specific domain for a given install`,
	Run:   installsDomainsDelete,
}

func init() {
	installsDomainsDeleteCmd.Flags().StringP("install", "i", "", "The install ID to delete the domain from")
	installsDomainsDeleteCmd.MarkFlagRequired("install")

	installsDomainsDeleteCmd.Flags().StringP("domain", "d", "", "The domain ID of the domain to be deleted")
	installsDomainsDeleteCmd.MarkFlagRequired("domain")
}

func installsDomainsDelete(cmd *cobra.Command, args []string) {
	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)

	installID, err := cmd.Flags().GetString("install")
	cobra.CheckErr(err)

	domainID, err := cmd.Flags().GetString("domain")
	cobra.CheckErr(err)

	format, err := cmd.Flags().GetString("format")
	cobra.CheckErr(err)

	result, err := api.InstallsDomainsDelete(installID, domainID)
	cobra.CheckErr(err)

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(result)
		cobra.CheckErr(err)

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return
	}

	fmt.Fprintf(os.Stdout, "%s deleted(%t) ", result.ID, result.IsDeleted)
}
