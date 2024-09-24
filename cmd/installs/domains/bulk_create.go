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

// installsDomainsBulkCreateCmd represents the installsDomainsDelete command
var installsDomainsBulkCreateCmd = &cobra.Command{
	Use:   "bulk-create",
	Short: "",
	Long:  ``,
	Run:   installsDomainsBulkCreate,
}

func init() {
	installsDomainsBulkCreateCmd.Flags().StringP("install", "i", "", "The install ID where you want to create the domains")
	installsDomainsBulkCreateCmd.MarkFlagRequired("install")

	installsDomainsBulkCreateCmd.Flags().StringP("body", "b", "", "The JSON of the domains to be created")
	installsDomainsBulkCreateCmd.MarkFlagRequired("body")

}

func installsDomainsBulkCreate(cmd *cobra.Command, args []string) {
	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	a := api.NewAPI(config)

	installID, err := cmd.Flags().GetString("install")
	cobra.CheckErr(err)

	body, err := cmd.Flags().GetString("body")
	cobra.CheckErr(err)

	// Using this to validate the JSON
	var bd api.BulkDomains
	err = json.NewDecoder(strings.NewReader(body)).Decode(&bd)
	if err != nil {
		cobra.CheckErr(err)
	}

	// TODO: validate domains
	if len(bd.Domains) == 0 {
		cobra.CheckErr(fmt.Errorf("no domains to create"))
	}

	if len(bd.Domains) > 20 {
		cobra.CheckErr(fmt.Errorf("max 20 domains can be created at once"))
	}

	format, err := cmd.Flags().GetString("format")
	cobra.CheckErr(err)

	result, err := a.InstallsDomainsBulkCreate(installID, body)
	cobra.CheckErr(err)

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(result)
		cobra.CheckErr(err)

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return
	}

	if len(result.Domains) > 0 {
		// NOTE: It's Redirect_S_ (with an S)
		for _, domains := range result.Domains {
			if domains.RedirectsTo.Name == "" {
				fmt.Fprintf(os.Stdout, "%s %s\n", domains.ID, domains.Name)
			} else {
				fmt.Fprintf(os.Stdout, "%s %s -> %s %s\n", domains.ID, domains.Name, domains.RedirectsTo.ID, domains.RedirectsTo.Name)
			}
		}
	} else {
		cmd.PrintErr("No domains created\n")
	}
}
