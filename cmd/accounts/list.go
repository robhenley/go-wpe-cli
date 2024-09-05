package accounts

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/robhenley/go-wpe-cli/cmd/types"
	"github.com/robhenley/go-wpe-cli/internal/api"
	"github.com/spf13/cobra"
)

// accountsCmd represents the accounts command
var accountsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List your WP Engine accounts",
	Long:  `Use this to list your WP Engine accounts.`,
	Run:   accountsList,
}

func init() {
	accountsListCmd.Flags().Int("page", 1, "The page to return")
	accountsListCmd.Flags().Int("limit", 100, "Limit the number of results")
}

func accountsList(cmd *cobra.Command, args []string) {
	page, err := cmd.Flags().GetInt("page")
	if err != nil {
		cmd.PrintErrf("Error: %s\n", err.Error())
		return
	}

	limit, err := cmd.Flags().GetInt("limit")
	if err != nil {
		cmd.PrintErrf("Error: %s\n", err.Error())
		return
	}

	format, err := cmd.Flags().GetString("format")
	if err != nil {
		cmd.PrintErrf("Error: %s\n", err.Error())
		return
	}

	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	config.Limit = limit

	api := api.NewAPI(config)
	accounts, err := api.AccountsList(page)
	if err != nil {
		cmd.PrintErrf("Error: %s\n", err.Error())
		return
	}

	if len(accounts) > 0 {
		if strings.ToLower(format) == "json" {
			j, err := json.Marshal(accounts)
			if err != nil {
				cmd.PrintErrf("Error: %s\n", err.Error())
				return
			}

			fmt.Fprintf(os.Stdout, "%s\n", j)
			return

		}

		for _, account := range accounts {
			fmt.Fprintf(os.Stdout, "%s %s\n", account.ID, account.Name)
		}
	}
}
