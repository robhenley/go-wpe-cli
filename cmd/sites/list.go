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
	Long:  `List the sites you have access to.`,
	Run:   sitesList,
}

func init() {
	listCmd.Flags().Int("page", 1, "The page to return")
	listCmd.Flags().Int("limit", 100, "Limit the number of results")
}

func sitesList(cmd *cobra.Command, args []string) {
	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)

	limit, err := cmd.Flags().GetInt("limit")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	config.Limit = limit

	api := api.NewAPI(config)

	page, err := cmd.Flags().GetInt("page")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	response := api.SitesList(page)

	format := cmd.Flag("format").Value.String()
	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(response)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return
	}

	for _, result := range response.Results {
		fmt.Fprintf(os.Stdout, "%s\t%s\t%s\n", result.ID, result.GroupName, result.Name)
	}

}
