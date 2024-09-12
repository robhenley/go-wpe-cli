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

// installsDomainsCreateCmd represents the installsDomainsGet command
var installsDomainsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Adds a new domain or redirect to an existing install",
	Long: `Adds a domain or redirect to a specific install and optionally sets
it as the primary domain.`,
	Run: installsDomainsCreate,
}

func init() {
	installsDomainsCreateCmd.Flags().String("install", "", "The install ID to operate on.")
	installsDomainsCreateCmd.MarkFlagRequired("install")

	installsDomainsCreateCmd.Flags().String("name", "", "The domain name to create (e.g. example.com)")
	installsDomainsCreateCmd.MarkFlagRequired("name")

	installsDomainsCreateCmd.Flags().String("redirect", "", "Redirect the domain to this domain")
	installsDomainsCreateCmd.Flags().Bool("primary", false, "Set the domain as a primary domain")
}

func installsDomainsCreate(cmd *cobra.Command, args []string) {
	installID, err := cmd.Flags().GetString("install")
	cobra.CheckErr(err)

	name, err := cmd.Flags().GetString("name")
	cobra.CheckErr(err)

	redirect, err := cmd.Flags().GetString("redirect")
	cobra.CheckErr(err)

	primary, err := cmd.Flags().GetBool("primary")
	cobra.CheckErr(err)

	format, err := cmd.Flags().GetString("format")
	cobra.CheckErr(err)

	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)

	domain, err := api.InstallsDomainsCreate(installID, name, redirect, primary)
	cobra.CheckErr(err)

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(domain)
		cobra.CheckErr(err)

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return
	}

	// NOTE: For /installs/{install_id}/domains it's redirectS_to not redirect_to
	fmt.Fprintf(os.Stdout, "%s %s %t %s", domain.ID, domain.Name, domain.Primary, domain.RedirectsTo.Name)
}
