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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your sites",
	Long: `List the sites you have access to. You can filter the results with
the filters flag.  An example usage is:

wpe sites list --filters "group=Clients,tag=Suspended,tag=Inactive"

NOTE: Both tags and groups are case insensitive.

`,
	Run: sitesList,
}

func init() {
	listCmd.Flags().StringSliceP("filters", "f", []string{}, "Filter the list of sites")
	listCmd.Flags().Int("page", 1, "The page to return")
	listCmd.Flags().Int("limit", 100, "Limit the number of results")
}

func sitesList(cmd *cobra.Command, args []string) {
	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)

	filters, err := cmd.Flags().GetStringSlice("filters")
	cobra.CheckErr(err)

	limit, err := cmd.Flags().GetInt("limit")
	cobra.CheckErr(err)
	config.Limit = limit

	format, err := cmd.Flags().GetString("format")
	cobra.CheckErr(err)

	page, err := cmd.Flags().GetInt("page")
	cobra.CheckErr(err)

	sites, err := api.SitesList(page, filters)
	cobra.CheckErr(err)

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(sites)
		cobra.CheckErr(err)

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return
	}

	if len(sites) > 0 {
		for _, site := range sites {
			fmt.Fprintf(os.Stdout, "%s\t%-15s\t%s\n", site.ID, site.GroupName, site.Name)
		}
	} else {
		fmt.Println("No sites were returned.")
	}

}
