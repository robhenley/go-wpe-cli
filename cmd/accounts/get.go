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

func init() {
	accountsGetCmd.Flags().StringP("account", "a", "", "The account ID to get the user from")
	accountsGetCmd.MarkFlagRequired("account")

}

func accountsGet(cmd *cobra.Command, args []string) {
	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)

	accountID, err := cmd.Flags().GetString("account")
	cobra.CheckErr(err)

	format, err := cmd.Flags().GetString("format")
	cobra.CheckErr(err)

	api := api.NewAPI(config)
	account, err := api.AccountsGet(accountID)
	cobra.CheckErr(err)

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(account)
		cobra.CheckErr(err)

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return

	}

	fmt.Fprintf(os.Stdout, "%s %s\n", account.ID, account.Name)
}
