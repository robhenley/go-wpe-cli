package sites

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/robhenley/go-wpe-cli/cmd/types"
	"github.com/robhenley/go-wpe-cli/internal/api"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new site",
	Long:  `Create a new site`,
	Run:   sitesCreate,
}

func init() {
	createCmd.Flags().StringP("account", "a", "", "The account ID to create the site under")
	createCmd.MarkFlagRequired("account")

	createCmd.Flags().StringP("name", "n", "", "The name of the site")
	createCmd.MarkFlagRequired("name")
}

func sitesCreate(cmd *cobra.Command, args []string) {
	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)

	accountID, err := cmd.Flags().GetString("account")
	cobra.CheckErr(err)

	name, err := cmd.Flags().GetString("name")
	cobra.CheckErr(err)

	format, err := cmd.Flags().GetString("format")
	cobra.CheckErr(err)

	site, err := api.SitesCreate(accountID, name)
	cobra.CheckErr(err)

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(site)
		cobra.CheckErr(err)

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return
	}

	fmt.Fprintf(os.Stdout, "%s %s\n", site.ID, site.Name)

}
