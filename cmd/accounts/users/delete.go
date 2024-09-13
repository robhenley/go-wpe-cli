package users

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/robhenley/go-wpe-cli/cmd/types"
	"github.com/robhenley/go-wpe-cli/internal/api"
	"github.com/spf13/cobra"
)

// accountsUsersDeleteCmd represents the accounts command
var accountsUsersDeleteCmd = &cobra.Command{
	Use:   "delete <account id> <user id>",
	Short: "Delete an account user",
	Long:  `This will remove the association this user has to this account. This delete is permanent and there is no confirmation prompt.`,
	Run:   accountsUsersDelete,
}

func init() {
	accountsUsersDeleteCmd.Flags().StringP("account", "a", "", "The account ID to delete the user from")
	accountsUsersDeleteCmd.MarkFlagRequired("account")

	accountsUsersDeleteCmd.Flags().StringP("user", "u", "", "The User ID to delete from the account")
	accountsUsersDeleteCmd.MarkFlagRequired("user")
}

func accountsUsersDelete(cmd *cobra.Command, args []string) {
	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)

	accountID, err := cmd.Flags().GetString("account")
	cobra.CheckErr(err)

	userID, err := cmd.Flags().GetString("user")
	cobra.CheckErr(err)

	format, err := cmd.Flags().GetString("format")
	cobra.CheckErr(err)

	result, err := api.AccountsUsersDelete(accountID, userID)
	cobra.CheckErr(err)

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(result)
		cobra.CheckErr(err)

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return

	}

	fmt.Fprintf(os.Stdout, "%s %t\n", result.ID, result.IsDeleted)
}
