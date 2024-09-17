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

// installsDomainsUpdateCmd represents the domains list command
var installsDomainsUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Set an existing domain as primary",
	Long:  `Sets a domain as the primary. Cannot set a duplicate, wildcard, or redirected domain as the primary.`,
	Run:   installsDomainsUpdate,
}

func init() {
	installsDomainsUpdateCmd.Flags().StringP("install", "i", "", "The install ID which has the domain you'd like to update")
	installsDomainsUpdateCmd.MarkFlagRequired("install")

	installsDomainsUpdateCmd.Flags().StringP("domain-id", "d", "", "The domain ID you'd like to set as primary")
	installsDomainsUpdateCmd.MarkFlagRequired("domain-id")

	installsDomainsUpdateCmd.Flags().StringP("redirect", "r", "", "The UUID of another domain record, or 'nil' to remove an existing redirect")
	installsDomainsUpdateCmd.Flags().Bool("primary", false, "Make the domain primary on the given install")
}

func installsDomainsUpdate(cmd *cobra.Command, args []string) {
	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)

	installID, err := cmd.Flags().GetString("install")
	cobra.CheckErr(err)

	domainID, err := cmd.Flags().GetString("domain-id")
	cobra.CheckErr(err)

	redirect, err := cmd.Flags().GetString("redirect")
	cobra.CheckErr(err)

	primary, err := cmd.Flags().GetBool("primary")
	cobra.CheckErr(err)

	format, err := cmd.Flags().GetString("format")
	cobra.CheckErr(err)

	response, err := api.InstallsDomainsUpdate(installID, domainID, redirect, primary)
	cobra.CheckErr(err)

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(response)
		cobra.CheckErr(err)

		fmt.Fprintf(os.Stdout, "%s", j)
		return
	}

	fmt.Fprintf(os.Stdout, "Domain %s updated\n", response.Name)

}
