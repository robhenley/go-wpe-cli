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

func sitesUpdate(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		fmt.Fprint(os.Stderr, "Error: Please provide a site id and a site name\n")
		cmd.Usage()
		return
	}

	siteID := args[0]
	siteName := args[1]

	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)

	api := api.NewAPI(config)

	site := api.SitesUpdate(siteID, siteName)

	format := cmd.Flag("format").Value.String()
	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(site)
		cobra.CheckErr(err)

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return
	}

	fmt.Fprintf(os.Stdout, "%s %s\n", site.ID, site.Name)

}
