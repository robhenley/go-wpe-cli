package installs

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
	Use:   "list [account id]",
	Short: "List your WordPress installations",
	Long:  `List your WordPress installations`,
	Run:   installsList,
}

func init() {
	listCmd.Flags().Int("page", 1, "The page to return")
	listCmd.Flags().Int("limit", 100, "Limit the number of results")
}

func installsList(cmd *cobra.Command, args []string) {
	var accountID string
	if len(args) == 1 {
		accountID = args[0]
	}

	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)

	page, err := cmd.Flags().GetInt("page")
	if err != nil {
		cmd.PrintErrf("Error: %v", err)
		return
	}

	limit, err := cmd.Flags().GetInt("limit")
	if err != nil {
		cmd.PrintErrf("Error: %v", err)
		return
	}
	config.Limit = limit

	api := api.NewAPI(config)
	installs, err := api.InstallsList(page, accountID)
	if err != nil {
		cmd.PrintErrf("Error: %v\n", err)
		return
	}

	format, err := cmd.Flags().GetString("format")
	if err != nil {
		cmd.PrintErrf("Error: %v", err)
		return
	}

	if strings.ToLower(format) == "json" {

		j, err := json.Marshal(installs.Results)
		if err != nil {
			cmd.PrintErrf("Error: %v\n", err)
			return
		}

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return
	}

	for _, install := range installs.Results {
		fmt.Fprintf(os.Stdout, "%s\t%-15s\t%-15s\t%s\n", install.ID, install.Environment, install.Name, install.PrimaryDomain)
	}

}
