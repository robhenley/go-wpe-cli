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

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a site",
	Long:  `Get a site by its site ID`,
	Run:   sitesGet,
}

func init() {
	getCmd.Flags().StringP("site", "s", "", "The ID of the site to get")
	getCmd.MarkFlagRequired("site")
}

func sitesGet(cmd *cobra.Command, args []string) {
	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)

	siteID, err := cmd.Flags().GetString("site")
	cobra.CheckErr(err)

	format, err := cmd.Flags().GetString("format")
	cobra.CheckErr(err)

	site, err := api.SitesGet(siteID)
	cobra.CheckErr(err)

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(site)
		cobra.CheckErr(err)

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return
	}

	fmt.Fprintf(os.Stdout, "%s\t%-15s\t%s\n", site.ID, site.GroupName, site.Name)

}
