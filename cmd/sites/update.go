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

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update <site id> <site name>",
	Short: "Update a site",
	Long:  `Long.`,
	Run:   sitesUpdate,
}

func initt() {
	updateCmd.Flags().StringP("site", "s", "", "The ID of the site to update")
	updateCmd.MarkFlagRequired("site")

	updateCmd.Flags().StringP("name", "n", "", "The site name")
	updateCmd.MarkFlagRequired("name")
}

func sitesUpdate(cmd *cobra.Command, args []string) {
	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)

	siteID, err := cmd.Flags().GetString("site")
	cobra.CheckErr(err)

	siteName, err := cmd.Flags().GetString("name")
	cobra.CheckErr(err)

	format, err := cmd.Flags().GetString("format")
	cobra.CheckErr(err)

	site := api.SitesUpdate(siteID, siteName)

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(site)
		cobra.CheckErr(err)

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return
	}

	fmt.Fprintf(os.Stdout, "%s %s\n", site.ID, site.Name)

}
