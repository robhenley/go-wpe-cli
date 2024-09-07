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

// accountsGetCmd represents the accounts command
var accountsGetCmd = &cobra.Command{
	Use:   "get <account id>",
	Short: "Get an account by ID",
	Long:  `Returns a single account`,
	Run:   accountsGet,
}

func accountsGet(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Usage()
		return
	}

	format, err := cmd.Flags().GetString("format")
	if err != nil {
		cmd.PrintErrf("Error: %s\n", err.Error())
		return
	}

	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)

	accountID := args[0]
	api := api.NewAPI(config)
	account, err := api.AccountsGet(accountID)
	if err != nil {
		cmd.PrintErrf("Error: %s\n", err.Error())
		return
	}

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(account)
		if err != nil {
			cmd.PrintErrf("Error: %s\n", err.Error())
			return
		}

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return

	}

	fmt.Fprintf(os.Stdout, "%s %s\n", account.ID, account.Name)
}