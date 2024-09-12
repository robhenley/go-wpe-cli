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
	Use:   "get <id>",
	Short: "Get a site",
	Long:  `Get a site by its site ID`,
	Run:   sitesGet,
}

func sitesGet(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Fprint(os.Stderr, "Error: Please provide a site id\n")
		cmd.Usage()
		return
	}
	id := args[0]

	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)

	api := api.NewAPI(config)
	site := api.SitesGet(id)

	format := cmd.Flag("format").Value.String()
	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(site)
		cobra.CheckErr(err)

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return
	}

	fmt.Fprintf(os.Stdout, "%s\t%-15s\t%s\n", site.ID, site.GroupName, site.Name)

}
